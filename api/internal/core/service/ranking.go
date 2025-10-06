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

func (rs *RankingService) RankingQuantity(
	ctx context.Context,
	event domain.EventID,
	state domain.StateID,
	mode domain.RankingMode,
) (int, error) {
	return rs.rr.RankingQuantity(ctx, event, state, mode)
}

func (rs *RankingService) Ranking(
	ctx context.Context,
	event domain.EventID,
	state domain.StateID,
	mode domain.RankingMode,
	lower domain.RankingLowerBound,
) ([]*domain.RankingEntry, error) {
	return rs.rr.Ranking(ctx, event, state, mode, lower)
}
