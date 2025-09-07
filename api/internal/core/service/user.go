package service

import (
	"context"

	"leinadium.dev/wca-ranking/internal/core/domain"
	"leinadium.dev/wca-ranking/internal/core/port"
)

type UserService struct {
	r port.UserRepository
}

func (s *UserService) User(ctx context.Context, id domain.WCAID) (*domain.User, error) {
	return s.r.User(ctx, id)
}

func (s *UserService) SetUser(ctx context.Context, user *domain.User) error {
	return s.r.SetUser(ctx, user)
}
