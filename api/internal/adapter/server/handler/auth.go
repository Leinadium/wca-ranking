package handler

import (
	"context"
	"fmt"

	"github.com/oapi-codegen/runtime/types"
	"leinadium.dev/wca-ranking/internal/adapter/server/schema"
	"leinadium.dev/wca-ranking/internal/core/domain"
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

func (s *ServerHandler) GetAuthCallback(
	ctx context.Context,
	request schema.GetAuthCallbackRequestObject,
) (schema.GetAuthCallbackResponseObject, error) {
	value, err := s.authService.ValidateCallback(ctx, request.Params.Code)
	if err != nil {
		return schema.GetAuthCallback500JSONResponse(ErrDefault), nil
	} else {
		return schema.GetAuthCallback200JSONResponse{
			AccessToken: &value.AccessToken,
			ExpiresAt:   &types.Date{Time: value.ExpiresAt},
		}, nil
	}
}

func (s *ServerHandler) PostPersonPersonId(
	ctx context.Context,
	request schema.PostPersonPersonIdRequestObject,
) (schema.PostPersonPersonIdResponseObject, error) {
	personId := domain.WCAID(request.PersonId)
	accessToken := request.Params.Authorization
	stateId := request.Body.StateId

	info, err := s.userService.ExternalUser(ctx, accessToken)
	if err != nil {
		return schema.PostPersonPersonId500JSONResponse(ErrWCADefault), nil
	}

	if info.WCAID != personId {
		return schema.PostPersonPersonId400JSONResponse(ErrUserInvalid), nil
	}

	wait, err := s.userService.HoursUntilAbleUpdate(ctx, personId)
	if err != nil {
		return schema.PostPersonPersonId500JSONResponse(ErrWCADefault), nil
	}
	if wait > 0 {
		msg := fmt.Sprintf(ErrCodeWCAUserWait, wait)
		return schema.PostPersonPersonId400JSONResponse{
			Code:    &ErrCodeWCAUserWait,
			Message: &msg,
		}, nil
	}

	if err := s.userService.SetUser(ctx, personId, *stateId); err != nil {
		return schema.PostPersonPersonId500JSONResponse(ErrDefault), nil
	}

	go s.syncService.Refresh(context.Background())

	return schema.PostPersonPersonId200Response{}, nil
}
