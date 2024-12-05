package repository

import (
	"go.opentelemetry.io/otel"

	"pkg/sql"
)

var tracer = otel.Tracer("/server/internal/services/checker/repository")

type CheckerRepository struct {
	db sql.SQL
}

func NewCheckerRepository(db sql.SQL, ) *CheckerRepository {
	return &CheckerRepository{
		db: db,
	}
}
