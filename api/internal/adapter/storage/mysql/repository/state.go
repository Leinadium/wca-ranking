package repository

import (
	"context"

	"leinadium.dev/wca-ranking/internal/adapter/storage/mysql/schema"
	"leinadium.dev/wca-ranking/internal/core/domain"
	"leinadium.dev/wca-ranking/pkg/utils"
)

func NewStateRepository(query *schema.Queries) *StateRepository {
	return &StateRepository{query: query}
}

type StateRepository struct {
	query *schema.Queries
}

func (s *StateRepository) States(ctx context.Context) ([]*domain.StateID, error) {
	rows, err := s.query.GetStates(ctx)
	if err != nil {
		return nil, err
	}
	return utils.Map(rows, func(row schema.AppState) *domain.StateID {
		return (*domain.StateID)(&row.StateID)
	}), nil
}
