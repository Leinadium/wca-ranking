package service

import (
	"context"
	"fmt"

	"leinadium.dev/wca-ranking/internal/adapter/config"
	"leinadium.dev/wca-ranking/internal/core/domain"
	"leinadium.dev/wca-ranking/internal/core/port"
)

type AuthService struct {
	requester port.AuthRequester
	config    *config.WCA
}

func NewAuthService(
	requester port.AuthRequester,
	config *config.WCA,
) *AuthService {
	return &AuthService{requester: requester, config: config}
}

func (s *AuthService) AuthEndpoint() (string, error) {
	return fmt.Sprintf(
		"%s?client_id=%s&redirect_uri=%s&response_type=code&scope=",
		s.config.Endpoints.OAuthAuthorize,
		s.config.ClientID,
		s.config.ClientSecret,
	), nil
}

func (s *AuthService) ValidateCallback(ctx context.Context, code string) (*domain.AuthPayload, error) {
	return s.requester.AccessToken(ctx, code)
}
