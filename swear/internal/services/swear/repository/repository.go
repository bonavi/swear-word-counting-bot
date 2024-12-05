package repository

import (
	"go.opentelemetry.io/otel"

	"pkg/sql"
)

var tracer = otel.Tracer("/server/internal/services/swear/repository")

type SwearRepository struct {
	db sql.SQL
}

func NewSwearRepository(db sql.SQL, ) *SwearRepository {
	return &SwearRepository{
		db: db,
	}
}
