package repository

import (
	"context"
	"time"

	"leinadium.dev/wca-ranking/internal/adapter/storage/mysql/schema"
)

func NewSyncRepository(query *schema.Queries) *SyncRepository {
	return &SyncRepository{query: query}
}

type SyncRepository struct {
	query *schema.Queries
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
