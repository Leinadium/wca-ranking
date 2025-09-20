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
	return nil
}

func (s *SyncService) Update() error {
	return nil
}

func (s *SyncService) Refresh() error {
	return nil
}

func (s *SyncService) CurrentDate(ctx context.Context) (*time.Time, error) {
	return s.r.CurrentDate(ctx)
}

func (s *SyncService) SetCurrentDate(time.Time) error {
	return nil
}
