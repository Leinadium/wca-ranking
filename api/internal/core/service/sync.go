package service

import (
	"context"
	"time"

	"leinadium.dev/wca-ranking/internal/core/domain"
	"leinadium.dev/wca-ranking/internal/core/port"
)

func NewSyncService(r port.SyncRepository) *SyncService {
	return &SyncService{r: r}
}

type SyncService struct {
	r port.SyncRepository
}

func (s *SyncService) ImportFile(file domain.File) error {
	return s.r.ImportFile(file)
}

func (s *SyncService) Update(ctx context.Context) error {
	return s.r.Update(ctx)
}

func (s *SyncService) Refresh(ctx context.Context) error {
	return s.r.Refresh(ctx)
}

func (s *SyncService) CurrentDate(ctx context.Context) (*time.Time, error) {
	return s.r.CurrentDate(ctx)
}

func (s *SyncService) SetCurrentDate(ctx context.Context, t time.Time) error {
	return s.r.SetCurrentDate(ctx, t)
}
