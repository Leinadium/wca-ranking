package handler

import (
	"context"

	"leinadium.dev/wca-ranking/internal/adapter/server/schema"
)

func (s *ServerHandler) GetAuthEndpoint(
	ctx context.Context,
	request schema.GetAuthEndpointRequestObject,
) (schema.GetAuthEndpointResponseObject, error) {
	value, err := s.authService.AuthEndpoint()
	if err != nil {
		return schema.GetAuthEndpoint500JSONResponse(ErrDefault), nil
	}
	return schema.GetAuthEndpoint200JSONResponse{
		Url: &value,
	}, nil
}
