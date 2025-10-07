package repository

import (
	"context"
	"database/sql"
	"os"
	"os/exec"
	"strconv"
	"time"

	"leinadium.dev/wca-ranking/internal/adapter/config"
	"leinadium.dev/wca-ranking/internal/adapter/storage/mysql"
	"leinadium.dev/wca-ranking/internal/adapter/storage/mysql/schema"
	"leinadium.dev/wca-ranking/internal/core/domain"
)

func NewSyncRepository(
	query *schema.Queries,
	config *config.DB,
	db *mysql.DB,
) *SyncRepository {
	return &SyncRepository{
		query:  query,
		config: config,
		db:     db,
	}
}

type SyncRepository struct {
	query  *schema.Queries
	config *config.DB
	db     *mysql.DB
}

func (s *SyncRepository) ImportFile(file domain.File) error {
	cmd := exec.Command(
		"/usr/bin/mariadb",
		"--host", s.config.Host,
		"--port", strconv.Itoa(s.config.Port),
		"--user", s.config.User,
		"-p"+s.config.Password,
		s.config.Tables.Dump,
		"-e", "source "+file.Name(),
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (s *SyncRepository) Update(ctx context.Context) error {
	tx, err := s.db.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	qtx := s.query.WithTx(tx)

	if err := qtx.RunUpdate(ctx); err != nil {
		return err
	}
	return qtx.RunRefresh(ctx)
}

func (s *SyncRepository) Refresh(ctx context.Context) error {
	return s.query.RunRefresh(ctx)
}

func (s *SyncRepository) CurrentDate(ctx context.Context) (*time.Time, error) {
	row, err := s.query.GetCurrentDate(ctx)
	if err != nil {
		return nil, err
	}
	if !row.Valid {
		return nil, nil
	}

	return &row.Time, nil
}

func (s *SyncRepository) SetCurrentDate(ctx context.Context, t time.Time) error {
	return s.query.SetCurrentDate(ctx, sql.NullTime{Time: t, Valid: true})
}
