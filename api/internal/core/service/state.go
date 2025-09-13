package service

import (
	"context"

	"leinadium.dev/wca-ranking/internal/core/domain"
	"leinadium.dev/wca-ranking/internal/core/port"
)

type StateService struct {
	sr port.StateRepository
}

func NewStateService(sr port.StateRepository) *StateService {
	return &StateService{sr: sr}
}

func (ss *StateService) States(ctx context.Context) ([]*domain.StateID, error) {
	return ss.sr.States(ctx)
}
