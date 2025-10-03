package handler

import (
	"context"

	"leinadium.dev/wca-ranking/internal/adapter/server/schema"
	"leinadium.dev/wca-ranking/internal/core/domain"
	"leinadium.dev/wca-ranking/pkg/utils"
)

func (s *ServerHandler) GetPersonResultsPersonId(
	ctx context.Context,
	request schema.GetPersonResultsPersonIdRequestObject,
) (schema.GetPersonResultsPersonIdResponseObject, error) {
	persons, err := s.personService.Results(ctx, domain.WCAID(request.PersonId))
	if err != nil {
		return schema.GetPersonResultsPersonId500JSONResponse{
			Code:    &ErrDefaultCode,
			Message: &ErrDefaultMessage,
		}, nil
	} else {
		return schema.GetPersonResultsPersonId200JSONResponse(
			utils.MapNotNull(persons, func(r *domain.PersonResult) schema.PersonResult {
				return schema.PersonResult{
					Event:          (*schema.Event)(&r.Event),
					Average:        utils.NullIntToPointer(r.Average),
					Single:         utils.NullIntToPointer(r.Single),
					RankingAverage: &r.RankingAverage,
					RankingSingle:  &r.RankingSingle,
				}
			}),
		), nil
	}
}

func (s *ServerHandler) GetPersonInfoPersonId(
	ctx context.Context,
	request schema.GetPersonInfoPersonIdRequestObject,
) (schema.GetPersonInfoPersonIdResponseObject, error) {
	person, err := s.personService.Person(ctx, domain.WCAID(request.PersonId))
	if err != nil {
		return schema.GetPersonInfoPersonId500JSONResponse{
			Code:    &ErrDefaultCode,
			Message: &ErrDefaultMessage,
		}, nil
	} else {
		return schema.GetPersonInfoPersonId200JSONResponse{
			Name:              &person.Name,
			State:             person.State.Ptr(),
			Registered:        &person.Registered,
			TotalCompetitions: &person.TotalCompetitions,
			StateCompetitions: utils.NullIntToPointer(person.StateCompetitions),
		}, nil
	}
}

func (s *ServerHandler) GetPersonModePersonId(
	ctx context.Context,
	request schema.GetPersonModePersonIdRequestObject,
) (schema.GetPersonModePersonIdResponseObject, error) {
	res, err := s.personService.Rankings(
		ctx,
		domain.WCAID(request.PersonId),
		domain.RankingMode(request.Mode),
	)
	if err != nil {
		return schema.GetPersonModePersonId500JSONResponse{
			Code:    &ErrDefaultCode,
			Message: &ErrDefaultMessage,
		}, nil
	} else {
		return schema.GetPersonModePersonId200JSONResponse(
			utils.MapNotNull(res, func(r *domain.PersonRanking) schema.PersonRanking {
				return schema.PersonRanking{
					Event:            (*schema.Event)(&r.Event),
					Ranking:          &r.Ranking,
					Best:             utils.NullIntToPointer(r.Best),
					CompetitionId:    r.CompetitionId.Ptr(),
					CompetitionName:  r.CompetitionName.Ptr(),
					CompetitionState: r.CompetitionState.Ptr(),
					Times:            utils.NullResultsToSlice(r.Times),
				}
			}),
		), nil
	}
}
