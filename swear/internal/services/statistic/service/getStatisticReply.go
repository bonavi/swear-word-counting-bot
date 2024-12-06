package service

import (
	"bytes"
	"context"

	"pkg/errors"
	"swearBot/internal/services/statistic/model"
)

func (s *StatisticService) GetDailyStatisticReply(ctx context.Context, req model.GetStatisticsReq) (string, error) {
	ctx, span := tracer.Start(ctx, "GetDailyStatisticReply")
	defer span.End()

	// Получаем статистику из базы данных
	statistics, err := s.statisticRepository.GetStatistics(ctx, req)
	if err != nil {
		return "", err
	}

	res := model.GetStatisticsRes{
		Date:       req.DateFrom,
		Statistics: statistics,
	}

	// Создаём буфер, в который будем записывать результат выполения шаблонизатора
	var buf bytes.Buffer

	// Выполняем шаблон
	if err = s.templates.dailyStatistic.Execute(&buf, res); err != nil {
		return "", errors.InternalServer.Wrap(err)
	}

	return buf.String(), nil
}
