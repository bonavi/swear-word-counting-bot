package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"server/internal/services/statistic/model"
	"server/internal/services/statistic/repository/statisticDDL"
)

func (r *StatisticRepository) SaveStatistic(ctx context.Context, req model.SaveStatisticsReq) error {
	ctx, span := tracer.Start(ctx, "SaveStatistic")
	defer span.End()

	q := sq.
		Insert(statisticDDL.Table).
		Columns(
			statisticDDL.ColumnMessageID,
			statisticDDL.ColumnChatID,
			statisticDDL.ColumnUserID,
			statisticDDL.ColumnSwear,
			statisticDDL.ColumnDatetime,
		)

	for _, swear := range req.Swears {
		q = q.Values(
			req.MessageID,
			req.ChatID,
			req.UserID,
			swear,
			req.Datetime,
		)
	}

	return r.db.Exec(ctx, q)
}
