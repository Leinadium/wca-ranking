package port

import (
	"context"

	"leinadium.dev/wca-ranking/internal/core/domain"
)

type SearchRepository interface {
	// Search querys the database searching for the name
	Search(ctx context.Context, query string) ([]*domain.SearchResult, error)
}

type SearchService interface {
	// Search querys the database searching for the name
	Search(ctx context.Context, query string) ([]*domain.SearchResult, error)
}
