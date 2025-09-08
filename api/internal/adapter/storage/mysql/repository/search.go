package repository

import (
	"context"

	"leinadium.dev/wca-ranking/internal/adapter/storage/mysql/schema"
	"leinadium.dev/wca-ranking/internal/core/domain"
	"leinadium.dev/wca-ranking/pkg/utils"
)

type SearchRepository struct {
	query *schema.Queries
}

func (s *SearchRepository) Search(ctx context.Context, query string) ([]*domain.SearchResult, error) {
	rows, err := s.query.Search(ctx, schema.SearchParams{Query: query})
	if err != nil {
		return nil, err
	}
	return utils.Map(rows, func(row schema.SearchRow) *domain.SearchResult {
		return &domain.SearchResult{
			WCAID:   domain.WCAID(row.WcaID),
			Name:    row.WcaName,
			StateID: domain.NullStateID(SQLNullString(row.StateID)),
		}
	}), nil
}
