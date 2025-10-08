package service

import (
	"context"
	"time"

	"leinadium.dev/wca-ranking/internal/adapter/config"
	"leinadium.dev/wca-ranking/internal/core/domain"
	"leinadium.dev/wca-ranking/internal/core/port"
)

type UserService struct {
	repository    port.UserRepository
	requester     port.UserRequester
	registerLimit time.Duration
}

func NewUserService(
	repository port.UserRepository,
	requester port.UserRequester,
	config *config.Auth,
) *UserService {
	return &UserService{
		repository:    repository,
		requester:     requester,
		registerLimit: time.Duration(config.RegisterTimeout) * time.Hour,
	}
}

func (s *UserService) ExternalUser(ctx context.Context, accessToken string) (*domain.UserBasic, error) {
	return s.requester.UserInfo(ctx, accessToken)
}

func (s *UserService) User(ctx context.Context, id domain.WCAID) (*domain.User, error) {
	return s.repository.User(ctx, id)
}

func (s *UserService) HoursUntilAbleUpdate(ctx context.Context, id domain.WCAID) (int, error) {
	if s.registerLimit <= 0 {
		return 0, nil
	}

	user, err := s.User(ctx, id)
	if err != nil {
		return 0, err
	}

	if time.Since(user.RegisterDate) < s.registerLimit {
		return int(time.Until(user.RegisterDate.Add(s.registerLimit)).Hours()), nil
	}
	return 0, nil
}

func (s *UserService) SetUser(ctx context.Context, id domain.WCAID, state string) error {
	return s.repository.SetUser(ctx, &domain.User{
		WCAID:        id,
		StateID:      state,
		RegisterDate: time.Now(),
	})
}
