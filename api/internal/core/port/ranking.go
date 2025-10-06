package port

import (
	"context"

	"leinadium.dev/wca-ranking/internal/core/domain"
)

type RankingRepository interface {
	// RankingQuantity get the number of total entries in the result
	RankingQuantity(
		context.Context,
		domain.EventID,
		domain.StateID,
		domain.RankingMode,
	) (
		int,
		error,
	)

	// Ranking gets the ranking for the event and state provided
	Ranking(
		ctx context.Context,
		event domain.EventID,
		state domain.StateID,
		mode domain.RankingMode,
		lower domain.RankingLowerBound,
	) (
		[]*domain.RankingEntry,
		error,
	)
}

type RankingService interface {
	// RankingQuantity get the number of total entries in the result
	RankingQuantity(
		context.Context,
		domain.EventID,
		domain.StateID,
		domain.RankingMode,
	) (
		int,
		error,
	)

	// Ranking gets the ranking for the event and state provided
	Ranking(
		context.Context,
		domain.EventID,
		domain.StateID,
		domain.RankingMode,
		domain.RankingLowerBound,
	) (
		[]*domain.RankingEntry,
		error,
	)
}
