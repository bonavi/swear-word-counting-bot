package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"pkg/ddlHelper"
	"server/internal/services/checker/repository/swearsDDL"
)

func (r *CheckerRepository) GetSwears(ctx context.Context) (swears []string, err error) {
	ctx, span := tracer.Start(ctx, "GetSwears")
	defer span.End()

	// Получаем маты
	rows, err := r.db.Query(ctx, sq.
		Select(ddlHelper.SelectAll).
		From(swearsDDL.Table),
	)
	if err != nil {
		return nil, err
	}

	// Проходимся по каждому мату
	for rows.Next() {

		var swear string
		if err = rows.Scan(&swear); err != nil {
			return nil, err
		}

		swears = append(swears, swear)
	}

	return swears, nil
}
