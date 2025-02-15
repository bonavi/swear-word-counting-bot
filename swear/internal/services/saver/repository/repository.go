package repository

import (
	"go.opentelemetry.io/otel"

	"pkg/sql"
)

var tracer = otel.Tracer("/server/internal/services/saver/repository") //nolint:unused

type SaverRepository struct {
	db sql.SQL
}

func NewSaverRepository(db sql.SQL, ) *SaverRepository {
	return &SaverRepository{
		db: db,
	}
}
