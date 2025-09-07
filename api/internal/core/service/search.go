package service

import (
	"context"

	"leinadium.dev/wca-ranking/internal/core/domain"
	"leinadium.dev/wca-ranking/internal/core/port"
)

type SearchService struct {
	ss port.SearchRepository
}

func (s *SearchService) Search(ctx context.Context, query string) ([]*domain.SearchResult, error) {
	return s.ss.Search(ctx, query)
}
