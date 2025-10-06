package repository

import (
	"context"
	"errors"

	"github.com/guregu/null"
	"leinadium.dev/wca-ranking/internal/adapter/storage/mysql/schema"
	"leinadium.dev/wca-ranking/internal/core/domain"
	"leinadium.dev/wca-ranking/pkg/utils"
)

type RankingRepository struct {
	query *schema.Queries
}

func NewRankingRepository(query *schema.Queries) *RankingRepository {
	return &RankingRepository{query: query}
}

func (r *RankingRepository) RankingQuantity(
	ctx context.Context,
	event domain.EventID,
	state domain.StateID,
	mode domain.RankingMode,
) (
	int, error,
) {
	var (
		q   int64
		err error
	)
	switch mode {
	case domain.RankingAverage:
		q, err = r.query.GetRankingAverageTotal(ctx, schema.GetRankingAverageTotalParams{
			Stateid: string(state),
			Eventid: string(event),
		})
	case domain.RankingSingle:
		q, err = r.query.GetRankingSingleTotal(ctx, schema.GetRankingSingleTotalParams{
			Stateid: string(state),
			Eventid: string(event),
		})
	}
	if err != nil {
		return 0, err
	}
	return int(q), nil
}

func (r *RankingRepository) Ranking(
	ctx context.Context,
	event domain.EventID,
	state domain.StateID,
	mode domain.RankingMode,
	lower domain.RankingLowerBound,
) (
	[]*domain.RankingEntry,
	error,
) {
	cursor := int32(lower)
	if cursor == 0 {
		cursor = -1
	}

	switch mode {
	case domain.RankingAverage:
		{
			rows, err := r.query.GetRankingAverage(ctx, schema.GetRankingAverageParams{
				Stateid:       string(state),
				Eventid:       string(event),
				Rankingcursor: cursor,
			})
			if err != nil {
				return nil, err
			}
			return utils.Map(rows, func(row schema.GetRankingAverageRow) *domain.RankingEntry {
				return &domain.RankingEntry{
					WCAID:            domain.WCAID(row.WcaID),
					Name:             row.Name,
					StateID:          domain.StateID(row.StateID),
					Best:             utils.SQLNullInt32(row.Best),
					Ranking:          int(row.Ranking),
					Registered:       row.Registered != 0,
					CompetitionID:    utils.SQLNullString(row.CompetitionID),
					CompetitionName:  utils.SQLNullString(row.CompetitionName),
					CompetitionState: utils.SQLNullString(row.CompetitionState),
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
			rows, err := r.query.GetRankingSingle(ctx, schema.GetRankingSingleParams{
				Stateid:       string(state),
				Eventid:       string(event),
				Rankingcursor: cursor,
			})
			if err != nil {
				return nil, err
			}
			return utils.Map(rows, func(row schema.GetRankingSingleRow) *domain.RankingEntry {
				return &domain.RankingEntry{
					WCAID:            domain.WCAID(row.WcaID),
					Name:             row.Name,
					StateID:          domain.StateID(row.StateID),
					Best:             utils.SQLNullInt32(row.Best),
					Ranking:          int(row.Ranking),
					Registered:       row.Registered != 0,
					CompetitionID:    utils.SQLNullString(row.CompetitionID),
					CompetitionName:  utils.SQLNullString(row.CompetitionName),
					CompetitionState: utils.SQLNullString(row.CompetitionState),
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
