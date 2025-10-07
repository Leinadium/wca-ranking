package port

import (
	"context"

	"leinadium.dev/wca-ranking/internal/core/domain"
)

type UserRequester interface {
	// UserInfo obtains the current basic information from WCA's API
	UserInfo(ctx context.Context, accessToken string) (*domain.UserBasic, error)
}

type UserRepository interface {
	// User gets the registered user in the database
	User(ctx context.Context, id domain.WCAID) (*domain.User, error)

	// SetUser updates or creates the user in the database
	SetUser(ctx context.Context, user *domain.User) error
}

type UserService interface {
	// UserInfo obtains the current basic information from WCA's API
	ExternalUser(ctx context.Context, accessToken string) (*domain.UserBasic, error)

	// User gets the registered user in the database
	User(ctx context.Context, id domain.WCAID) (*domain.User, error)

	// SetUser updates or creates the user in the database
	SetUser(ctx context.Context, user *domain.User) error
}
