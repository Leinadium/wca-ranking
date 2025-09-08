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

func (r *RankingRepository) Ranking(
	ctx context.Context,
	event domain.EventID,
	state domain.StateID,
	mode domain.RankingMode,
) (
	[]*domain.RankingEntry,
	error,
) {
	switch mode {
	case domain.RankingAverage:
		{
			rows, err := r.query.GetRankingAverage(ctx, schema.GetRankingAverageParams{
				Stateid: string(state),
				Eventid: string(event),
			})
			if err != nil {
				return nil, err
			}
			return utils.Map(rows, func(row schema.GetRankingAverageRow) *domain.RankingEntry {
				return &domain.RankingEntry{
					WCAID:            domain.WCAID(row.WcaID),
					Name:             row.Name,
					StateID:          domain.StateID(row.StateID),
					Best:             SQLNullInt32(row.Best),
					Ranking:          int(row.Ranking),
					Registered:       row.Registered != 0,
					CompetitionID:    SQLNullString(row.CompetitionID),
					CompetitionName:  SQLNullString(row.CompetitionName),
					CompetitionState: SQLNullString(row.CompetitionState),
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
	case domain.RankingSingle:
		{
			rows, err := r.query.GetRankingSingle(ctx, schema.GetRankingSingleParams{
				Stateid: string(state),
				Eventid: string(event),
			})
			if err != nil {
				return nil, err
			}
			return utils.Map(rows, func(row schema.GetRankingSingleRow) *domain.RankingEntry {
				return &domain.RankingEntry{
					WCAID:            domain.WCAID(row.WcaID),
					Name:             row.Name,
					StateID:          domain.StateID(row.StateID),
					Best:             SQLNullInt32(row.Best),
					Ranking:          int(row.Ranking),
					Registered:       row.Registered != 0,
					CompetitionID:    SQLNullString(row.CompetitionID),
					CompetitionName:  SQLNullString(row.CompetitionName),
					CompetitionState: SQLNullString(row.CompetitionState),
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
	default:
		return nil, errors.New("invalid ranking mode")
	}
}
