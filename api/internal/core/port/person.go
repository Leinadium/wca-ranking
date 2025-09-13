package port

import (
	"context"

	"leinadium.dev/wca-ranking/internal/core/domain"
)

type PersonRepository interface {
	// Person gets the basic information about a person in the database
	Person(ctx context.Context, id domain.WCAID) (*domain.Person, error)

	// Results gets the current best results and its current standings of the person
	Results(ctx context.Context, id domain.WCAID) ([]*domain.PersonResult, error)

	// Rankings gets the details of rankings for all events using the provided mode for a person
	Rankings(ctx context.Context, id domain.WCAID, mode domain.RankingMode) ([]*domain.PersonRanking, error)
}

type PersonService interface {
	// Person gets the basic information about a person in the database
	Person(ctx context.Context, id domain.WCAID) (*domain.Person, error)

	// Results gets the current best results and its current standings of the person
	Results(ctx context.Context, id domain.WCAID) ([]*domain.PersonResult, error)

	// Rankings gets the details of rankings for all events using the provided mode for a person
	Rankings(ctx context.Context, id domain.WCAID, mode domain.RankingMode) ([]*domain.PersonRanking, error)
}
