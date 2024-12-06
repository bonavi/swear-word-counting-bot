package repository

import (
	"go.opentelemetry.io/otel"

	"pkg/sql"
)

var tracer = otel.Tracer("/server/internal/services/statistic/repository")

type StatisticRepository struct {
	db sql.SQL
}

func NewStatisticRepository(db sql.SQL, ) *StatisticRepository {
	return &StatisticRepository{
		db: db,
	}
}
