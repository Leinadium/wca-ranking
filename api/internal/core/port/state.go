package port

import (
	"context"

	"leinadium.dev/wca-ranking/internal/core/domain"
)

type StateRepository interface {
	// States gets all the states available for querying
	States(ctx context.Context) ([]*domain.StateID, error)
}

type StateService interface {
	// States gets all the states available for querying
	States(ctx context.Context) ([]*domain.StateID, error)
}
