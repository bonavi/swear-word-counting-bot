package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"swearBot/internal/ddl/swearsDDL"
	"swearBot/internal/services/swear/model"
)

func (r *SwearRepository) AddSwears(ctx context.Context, req model.AddSwearsReq) (err error) {
	ctx, span := tracer.Start(ctx, "AddSwears")
	defer span.End()

	q := sq.
		Insert(swearsDDL.Table).
		Columns(
			swearsDDL.ColumnText,
			swearsDDL.ColumnDatetime,
			swearsDDL.ColumnUserID,
		)

	// Получаем маты
	for _, swear := range req.Swears {
		q = q.Values(
			swear,
			req.Datetime,
			req.UserID,
		)
	}

	return r.db.Exec(ctx, q)
}
