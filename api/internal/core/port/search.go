package port

import (
	"context"

	"leinadium.dev/wca-ranking/internal/core/domain"
)

type SearchService interface {
	// Search querys the database searching for the name
	Search(ctx context.Context, query string) ([]*domain.SearchResult, error)
}
