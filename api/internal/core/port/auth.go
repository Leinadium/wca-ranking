package port

import (
	"context"

	"leinadium.dev/wca-ranking/internal/core/domain"
)

type AuthRequester interface {
	// AccessToken obtains an access token from a code auth
	AccessToken(ctx context.Context, code string) (*domain.AuthPayload, error)
}

type AuthService interface {
	// AuthEndpoint generates an authentication endpoint to be used by the user
	AuthEndpoint() (string, error)

	// ValidateCallback validates the user from a oauth callback code
	ValidateCallback(ctx context.Context, code string) (*domain.AuthPayload, error)
}
