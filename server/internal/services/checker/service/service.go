package service

import (
	"context"

	"go.opentelemetry.io/otel"

	"server/internal/services/checker/model"
	checkerRepository "server/internal/services/checker/repository"
)

var tracer = otel.Tracer("/server/internal/services/checker/service")

type CheckerService struct {
	checkerRepository CheckerRepository
}

var _ CheckerRepository = new(checkerRepository.CheckerRepository)

type CheckerRepository interface {
	GetSwears(context.Context) ([]string, error)
	SaveStatistic(context.Context, model.SaveStatisticsReq) error
}

func NewCheckerService(
	checkerRepository CheckerRepository,
) *CheckerService {
	return &CheckerService{
		checkerRepository: checkerRepository,
	}
}
