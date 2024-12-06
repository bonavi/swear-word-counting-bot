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
	templates           *Templates
}

var _ StatisticRepository = new(statisticRrpository.StatisticRepository)

type StatisticRepository interface {
	SaveStatistic(context.Context, model.SaveStatisticsReq) error
	GetStatistics(ctx context.Context, req model.GetStatisticsReq) ([]model.Statistic, error)
}

func NewStatisticService(
	statisticRepository StatisticRepository,
) (*StatisticService, error) {

	templates, err := NewTemplates()
	if err != nil {
		return nil, err
	}

	return &StatisticService{
		statisticRepository: statisticRepository,
		templates:           templates,
	}, nil
}
