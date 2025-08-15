package port

import (
	"context"

	"leinadium.dev/wca-ranking/internal/core/domain"
)

type WCATokenService interface {
	// AuthEndpoint generates an authentication endpoint to be used by the user
	AuthEndpoint() (string, error)

	// ValidateCallback validates the user from a oauth callback code
	ValidateCallback(ctx context.Context, code string) (*domain.WCATokenPayload, error)
}
