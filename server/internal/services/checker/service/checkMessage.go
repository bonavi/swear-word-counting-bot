package service

import (
	"context"

	"pkg/log"
	"server/internal/services/checker/model"
	"server/internal/services/checker/service/utils"
	statisticModel "server/internal/services/statistic/model"
)

func (s *CheckerService) CheckMessage(ctx context.Context, req model.CheckMessageReq) error {
	ctx, span := tracer.Start(ctx, "CheckMessage")
	defer span.End()

	// Получаем все маты
	swearsMap, err := s.swearService.GetSwears(ctx)
	if err != nil {
		return err
	}

	// Проверяем на маты
	swears := utils.GetSwears(req.Message.Text, swearsMap)

	// Если матов нет, выходим
	if len(swears) == 0 {
		return nil
	}

	log.Info(ctx, "Матершиник найден")

	// Сохраняем статистику
	if err = s.statisticRepository.SaveStatistic(ctx, statisticModel.SaveStatisticsReq{
		UserID:    req.User.ID,
		MessageID: req.Message.ID,
		ChatID:    req.Message.ChatID,
		Swears:    swears,
		Datetime:  req.Message.Datetime,
	}); err != nil {
		return err
	}

	return nil
}
