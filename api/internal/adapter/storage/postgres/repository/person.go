package repository

import (
	"context"

	"github.com/guregu/null"
	"leinadium.dev/wca-ranking/internal/adapter/storage/postgres/schema"
	"leinadium.dev/wca-ranking/internal/core/domain"
	"leinadium.dev/wca-ranking/pkg/utils"
)

type PersonRepository struct {
	query *schema.Queries
}

func (p *PersonRepository) Person(ctx context.Context, id domain.WCAID) (*domain.Person, error) {
	row, err := p.query.GetPersonInfo(ctx, schema.GetPersonInfoParams{
		Wcaid: string(id),
	})
	if err != nil {
		return nil, err
	}
	return &domain.Person{
		Name:              row.Name,
		State:             SQLNullString(row.StateID),
		Registered:        row.Registered != 0,
		TotalCompetitions: int(row.TotalCompetitions),
		StateCompetitions: SQLNullInt32(row.StateCompetitions),
	}, nil
}

func (p *PersonRepository) Results(ctx context.Context, id domain.WCAID) ([]*domain.PersonResult, error) {
	rows, err := p.query.GetPersonResults(ctx, string(id))
	if err != nil {
		return nil, err
	}

	return utils.Map(rows, func(row schema.GetPersonResultsRow) *domain.PersonResult {
		return &domain.PersonResult{
			Event:          row.Event,
			Single:         null.IntFrom(int64(row.Single)),
			Average:        null.IntFrom(int64(row.Average)),
			RankingSingle:  int(row.RankingSingle),
			RankingAverage: int(row.RankingAverage),
		}
	}), nil
}

func (p *PersonRepository) Rankings(ctx context.Context, id domain.WCAID, mode domain.RankingMode) ([]*domain.PersonRanking, error) {
	if mode == domain.RankingAverage {
		rows, err := p.query.GetPersonAverage(ctx, string(id))
		if err != nil {
			return nil, err
		}
		return utils.Map(rows, func(row schema.GetPersonAverageRow) *domain.PersonRanking {
			return &domain.PersonRanking{
				Mode:             mode,
				Event:            row.EventID,
				Ranking:          int(row.Ranking),
				Best:             SQLNullInt32(row.Best),
				CompetitionId:    SQLNullString(row.CompetitionID),
				CompetitionName:  SQLNullString(row.CompetitionID),
				CompetitionState: SQLNullString(row.CompetitionID),
				Times: [5]null.Int{
					SQLNullInt32(row.Time1),
					SQLNullInt32(row.Time2),
					SQLNullInt32(row.Time3),
					SQLNullInt32(row.Time4),
					SQLNullInt32(row.Time5),
				},
			}
		}), nil

	}

	return nil, nil
}
