package port

import (
	"context"

	"leinadium.dev/wca-ranking/internal/core/domain"
)

type RankingRepository interface {
	// Ranking gets the ranking for the event and state provided
	Ranking(ctx context.Context, event domain.EventID, state domain.StateID, mode domain.RankingMode) ([]*domain.RankingEntry, error)
}

type RankingService interface {
	// Ranking gets the ranking for the event and state provided
	Ranking(ctx context.Context, event domain.EventID, state domain.StateID, mode domain.RankingMode) ([]*domain.RankingEntry, error)
}
