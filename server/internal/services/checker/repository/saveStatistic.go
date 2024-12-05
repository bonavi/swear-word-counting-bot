package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"server/internal/services/checker/model"
	"server/internal/services/checker/repository/statisticDDL"
)

func (r *CheckerRepository) SaveStatistic(ctx context.Context, req model.SaveStatisticsReq) error {
	ctx, span := tracer.Start(ctx, "SaveStatistic")
	defer span.End()

	q := sq.
		Insert(statisticDDL.Table).
		Columns(
			statisticDDL.ColumnMessageID,
			statisticDDL.ColumnChatID,
			statisticDDL.ColumnUserID,
			statisticDDL.ColumnSwear,
		)

	for _, swear := range req.Swears {
		q = q.Values(req.MessageID, req.ChatID, req.UserID, swear)
	}

	return r.db.Exec(ctx, q)
}
