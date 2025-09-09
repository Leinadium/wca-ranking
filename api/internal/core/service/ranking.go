package service

import (
	"context"

	"leinadium.dev/wca-ranking/internal/core/domain"
	"leinadium.dev/wca-ranking/internal/core/port"
)

type RankingService struct {
	rr port.RankingService
}

func NewRankingService(rr port.RankingRepository) *RankingService {
	return &RankingService{rr: rr}
}

func (rs *RankingService) Ranking(
	ctx context.Context,
	event domain.EventID,
	state domain.StateID,
	mode domain.RankingMode,
) ([]*domain.RankingEntry, error) {
	return rs.rr.Ranking(ctx, event, state, mode)
}
