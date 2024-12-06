package service

import (
	"context"
	"strings"

	"pkg/log"
	"swearBot/internal/services/checker/model"
	"swearBot/internal/services/checker/service/utils"
	statisticModel "swearBot/internal/services/statistic/model"
)

var replyTexts = map[string]string{
	"пошел нахуй": "А не пошел бы ты сам нахуй, козел",
	"пошла нахуй": "А не пошел бы ты сам нахуй, козел",
}

func (s *CheckerService) CheckMessage(ctx context.Context, req model.CheckMessageReq) (string, error) {
	ctx, span := tracer.Start(ctx, "CheckMessage")
	defer span.End()

	// Кастомный хэндлинг
	if reply, ok := replyTexts[strings.ToLower(req.Message.Text)]; ok {
		return reply, nil
	}

	// Получаем все маты
	swearsMap, err := s.swearService.GetSwears(ctx)
	if err != nil {
		return "", err
	}

	// Проверяем на маты
	swears := utils.GetSwears(req.Message.Text, swearsMap)

	// Если матов нет, выходим
	if len(swears) == 0 {
		return "", nil
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
		return "", err
	}

	return "", nil
}
