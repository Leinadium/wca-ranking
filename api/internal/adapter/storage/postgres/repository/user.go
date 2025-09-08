package repository

import (
	"context"

	"leinadium.dev/wca-ranking/internal/adapter/storage/postgres/schema"
	"leinadium.dev/wca-ranking/internal/core/domain"
)

type UserRepository struct {
	query *schema.Queries
}

func (r *UserRepository) User(ctx context.Context, id domain.WCAID) (*domain.User, error) {
	row, err := r.query.GetUser(ctx, string(id))
	if err != nil {
		return nil, err
	}
	return &domain.User{
		WcaID:        row.WcaID,
		StateID:      row.StateID,
		RegisterDate: row.RegisterDate,
	}, nil
}

func (r *UserRepository) SetUser(ctx context.Context, user *domain.User) error {
	return r.query.SetUser(ctx, schema.SetUserParams{
		Wcaid:        user.WcaID,
		Stateid:      user.StateID,
		Registerdate: user.RegisterDate,
	})
}
