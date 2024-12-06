package service

import (
	"context"

	"go.opentelemetry.io/otel"

	"swearBot/internal/services/statistic/model"
	statisticRrpository "swearBot/internal/services/statistic/repository"
)

var tracer = otel.Tracer("/server/internal/services/statistic/service")

type StatisticService struct {
	statisticRepository StatisticRepository
}

var _ StatisticRepository = new(statisticRrpository.StatisticRepository)

type StatisticRepository interface {
	SaveStatistic(context.Context, model.SaveStatisticsReq) error
}

func NewStatisticService(
	statisticRepository StatisticRepository,
) *StatisticService {
	return &StatisticService{
		statisticRepository: statisticRepository,
	}
}
