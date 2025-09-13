package handler

import (
	"context"

	"leinadium.dev/wca-ranking/internal/adapter/server/schema"
	"leinadium.dev/wca-ranking/internal/core/domain"
	"leinadium.dev/wca-ranking/pkg/utils"
)

func (s *ServerHandler) GetStates(ctx context.Context, request schema.GetStatesRequestObject) (schema.GetStatesResponseObject, error) {
	states, err := s.stateService.States(ctx)
	if err != nil {
		return schema.GetStates500JSONResponse{
			Code:    &ErrDefaultCode,
			Message: &ErrDefaultMessage,
		}, nil
	} else {
		return schema.GetStates200JSONResponse(
			utils.MapNotNull(states, func(state *domain.StateID) string {
				return string(*state)
			}),
		), nil
	}
}
