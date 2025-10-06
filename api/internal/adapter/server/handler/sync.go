package handler

import (
	"context"

	"leinadium.dev/wca-ranking/internal/adapter/server/schema"
)

func (s *ServerHandler) GetStatus(ctx context.Context, request schema.GetStatusRequestObject) (schema.GetStatusResponseObject, error) {
	status, err := s.syncService.CurrentDate(ctx)
	if err != nil {
		return schema.GetStatus500JSONResponse(ErrDefault), nil
	} else {
		return schema.GetStatus200JSONResponse{
			LastUpdate: status,
		}, nil
	}
}
