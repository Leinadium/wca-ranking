package repository

import (
	"context"
	"errors"

	"github.com/guregu/null"
	"leinadium.dev/wca-ranking/internal/adapter/storage/mysql/schema"
	"leinadium.dev/wca-ranking/internal/core/domain"
	"leinadium.dev/wca-ranking/pkg/utils"
)

type PersonRepository struct {
	query *schema.Queries
}

func NewPersonRepository(query *schema.Queries) *PersonRepository {
	return &PersonRepository{query: query}
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
		State:             utils.SQLNullString(row.StateID),
		Registered:        row.Registered != 0,
		TotalCompetitions: int(row.TotalCompetitions),
		StateCompetitions: utils.SQLNullInt32(row.StateCompetitions),
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
	switch mode {
	case domain.RankingAverage:
		{
			rows, err := p.query.GetPersonAverage(ctx, string(id))
			if err != nil {
				return nil, err
			}
			return utils.Map(rows, func(row schema.GetPersonAverageRow) *domain.PersonRanking {
				return &domain.PersonRanking{
					Mode:             mode,
					Event:            row.EventID,
					Ranking:          int(row.Ranking),
					Best:             utils.SQLNullInt32(row.Best),
					CompetitionId:    utils.SQLNullString(row.CompetitionID),
					CompetitionName:  utils.SQLNullString(row.CompetitionID),
					CompetitionState: utils.SQLNullString(row.CompetitionID),
					Times: [5]null.Int{
						utils.SQLNullInt32(row.Time1),
						utils.SQLNullInt32(row.Time2),
						utils.SQLNullInt32(row.Time3),
						utils.SQLNullInt32(row.Time4),
						utils.SQLNullInt32(row.Time5),
					},
				}
			}), nil
		}
	case domain.RankingSingle:
		{
			rows, err := p.query.GetPersonSingle(ctx, string(id))
			if err != nil {
				return nil, err
			}
			return utils.Map(rows, func(row schema.GetPersonSingleRow) *domain.PersonRanking {
				return &domain.PersonRanking{
					Mode:             mode,
					Event:            row.EventID,
					Ranking:          int(row.Ranking),
					Best:             utils.SQLNullInt32(row.Best),
					CompetitionId:    utils.SQLNullString(row.CompetitionID),
					CompetitionName:  utils.SQLNullString(row.CompetitionID),
					CompetitionState: utils.SQLNullString(row.CompetitionID),
					Times: [5]null.Int{
						utils.SQLNullInt32(row.Time1),
						utils.SQLNullInt32(row.Time2),
						utils.SQLNullInt32(row.Time3),
						utils.SQLNullInt32(row.Time4),
						utils.SQLNullInt32(row.Time5),
					},
				}
			}), nil
		}
	default:
		return nil, errors.New("invalid ranking mode")
	}
}
