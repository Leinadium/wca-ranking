package handler

import (
	"context"

	"leinadium.dev/wca-ranking/internal/adapter/server/schema"
	"leinadium.dev/wca-ranking/internal/core/domain"
	"leinadium.dev/wca-ranking/pkg/utils"
)

func (s *ServerHandler) GetRankingModeEventState(
	ctx context.Context,
	request schema.GetRankingModeEventStateRequestObject,
) (
	schema.GetRankingModeEventStateResponseObject,
	error,
) {
	q, err := s.rankingService.RankingQuantity(
		ctx,
		domain.EventID(request.Event),
		domain.StateID(request.State),
		domain.RankingMode(request.Mode),
	)
	if err != nil {
		return schema.GetRankingModeEventState500JSONResponse(ErrDefault), nil
	}

	var lower int
	if request.Params.From == nil {
		lower = 0
	} else {
		lower = *request.Params.From
	}

	results, err := s.rankingService.Ranking(
		ctx,
		domain.EventID(request.Event),
		domain.StateID(request.State),
		domain.RankingMode(request.Mode),
		domain.RankingLowerBound(lower),
	)
	if err != nil {
		return schema.GetRankingModeEventState500JSONResponse(ErrDefault), nil
	}

	var next *int
	last, ok := utils.Last(results)
	if !ok {
		next = nil
	} else {
		next = &last.Ranking
	}

	res := utils.MapNotNull(results, func(r *domain.RankingEntry) schema.RankingEntry {
		return schema.RankingEntry{
			WcaId:            (*schema.WCAID)(&r.WCAID),
			StateId:          (*string)(&r.StateID),
			Name:             &r.Name,
			Ranking:          &r.Ranking,
			Registered:       &r.Registered,
			Best:             utils.NullIntToPointer(r.Best),
			CompetitionId:    utils.NullStringtoPointer(r.CompetitionID),
			CompetitionName:  utils.NullStringtoPointer(r.CompetitionName),
			CompetitionState: utils.NullStringtoPointer(r.CompetitionState),
		}
	})

	return schema.GetRankingModeEventState200JSONResponse{
		TotalItems: &q,
		NextItem:   next,
		Results:    &res,
	}, nil
}
