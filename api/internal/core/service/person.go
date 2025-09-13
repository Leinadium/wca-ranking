package service

import (
	"context"

	"leinadium.dev/wca-ranking/internal/core/domain"
	"leinadium.dev/wca-ranking/internal/core/port"
)

type PersonService struct {
	pr port.PersonRepository
}

func NewPersonService(pr port.PersonRepository) *PersonService {
	return &PersonService{pr: pr}
}

func (ps *PersonService) Person(ctx context.Context, id domain.WCAID) (*domain.Person, error) {
	return ps.pr.Person(ctx, id)
}

func (ps *PersonService) Results(ctx context.Context, id domain.WCAID) ([]*domain.PersonResult, error) {
	return ps.pr.Results(ctx, id)
}

func (ps *PersonService) Rankings(ctx context.Context, id domain.WCAID, mode domain.RankingMode) ([]*domain.PersonRanking, error) {
	return ps.pr.Rankings(ctx, id, mode)
}
