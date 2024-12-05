package service

import (
	"context"

	"pkg/log"
	"pkg/slices"
	"server/internal/services/checker/model"
	"server/internal/services/checker/service/utils"
)

func (s *CheckerService) CheckMessage(ctx context.Context, req model.CheckMessageReq) error {
	ctx, span := tracer.Start(ctx, "CheckMessage")
	defer span.End()

	// Получаем все маты
	swearsDictionary, err := s.checkerRepository.GetSwears(ctx)
	if err != nil {
		return err
	}

	// Делаем мапу
	swearsMap := slices.GetMapValueStruct(swearsDictionary, func(swear string) string { return swear })

	// Проверяем на маты
	swears := utils.GetSwears(req.Message.Text, swearsMap)

	// Если матов нет, выходим
	if len(swears) == 0 {
		return nil
	}

	log.Info(ctx, "Матершиник найден")

	// Сохраняем статистику
	if err = s.checkerRepository.SaveStatistic(ctx, model.SaveStatisticsReq{
		UserID:    req.User.ID,
		MessageID: req.Message.ID,
		ChatID:    req.Message.ChatID,
		Swears:    swears,
	}); err != nil {
		return err
	}

	return nil
}
