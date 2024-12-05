package service

import (
	"context"

	"go.opentelemetry.io/otel"

	"swearBot/internal/services/swear/model"
	swearRepository "swearBot/internal/services/swear/repository"
)

var tracer = otel.Tracer("/server/internal/services/swear/service")

type SwearService struct {
	swearRepository SwearRepository
}

var _ SwearRepository = new(swearRepository.SwearRepository)

type SwearRepository interface {
	AddSwears(context.Context, model.AddSwearsReq) error
	GetSwears(context.Context) ([]string, error)
}

func NewSwearService(
	swearRepository SwearRepository,
) *SwearService {
	return &SwearService{
		swearRepository: swearRepository,
	}
}
