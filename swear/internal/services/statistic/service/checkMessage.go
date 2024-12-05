package service

import (
	"context"

	"swearBot/internal/services/statistic/model"
)

func (s *StatisticService) SaveStatistic(ctx context.Context, req model.SaveStatisticsReq) error {
	ctx, span := tracer.Start(ctx, "SaveStatistic")
	defer span.End()

	return s.statisticRepository.SaveStatistic(ctx, req)
}
