package port

import (
	"context"

	"leinadium.dev/wca-ranking/internal/core/domain"
)

type UserRepository interface {
	// User gets the registered user in the database
	User(ctx context.Context, id domain.WCAID) (*domain.User, error)

	// SetUser updates or creates the user in the database
	SetUser(ctx context.Context, user *domain.User) error
}

type UserService interface {
	// User gets the registered user in the database
	User(ctx context.Context, id domain.WCAID) (*domain.User, error)

	// SetUser updates or creates the user in the database
	SetUser(ctx context.Context, user *domain.User) error
}
