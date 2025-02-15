package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"pkg/errors"
	"swearBot/internal/ddl/swearsDDL"
)

func (r *SwearRepository) GetSwears(ctx context.Context) (swears []string, err error) {
	ctx, span := tracer.Start(ctx, "GetSwears")
	defer span.End()

	// Получаем маты
	rows, err := r.db.Query(ctx, sq.
		Select(swearsDDL.ColumnText).
		From(swearsDDL.Table),
	)
	if err != nil {
		return nil, err
	}

	// Проходимся по каждому мату
	for rows.Next() {

		var swear string
		if err = rows.Scan(&swear); err != nil {
			return nil, errors.InternalServer.Wrap(err)
		}

		swears = append(swears, swear)
	}

	return swears, nil
}
