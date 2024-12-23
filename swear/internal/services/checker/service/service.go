package service

import (
	"context"

	"go.opentelemetry.io/otel"

	statisticModel "swearBot/internal/services/statistic/model"
	statisticRepository "swearBot/internal/services/statistic/repository"
	swearService "swearBot/internal/services/swear/service"
)

var tracer = otel.Tracer("/server/internal/services/checker/service")

type CheckerService struct {
	statisticRepository StatisticRepository
	swearService        SwearService
}

var _ StatisticRepository = new(statisticRepository.StatisticRepository)

type StatisticRepository interface {
	SaveStatistic(context.Context, statisticModel.SaveStatisticsReq) error
}

var _ SwearService = new(swearService.SwearService)

type SwearService interface {
	GetSwears(context.Context) (map[string]struct{}, error)
}

func NewCheckerService(
	checkerRepository StatisticRepository,
	swearService SwearService,
) *CheckerService {
	return &CheckerService{
		statisticRepository: checkerRepository,
		swearService:        swearService,
	}
}
