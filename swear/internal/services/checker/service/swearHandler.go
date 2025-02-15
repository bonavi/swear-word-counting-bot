package service

import (
	"context"

	"gopkg.in/telebot.v3"

	"pkg/errors"
	"swearBot/internal/services/checker/service/utils"
	statisticModel "swearBot/internal/services/statistic/model"
)

func (s *CheckerService) swearHandler(ctx context.Context, message *telebot.Message, user *telebot.User) error {

	// Получаем все маты из бд
	swearsMap, err := s.swearService.GetSwears(ctx)
	if err != nil {
		return err
	}

	// Проверяем на маты
	swears := utils.GetSwears(message.Text, swearsMap)

	// Если матов нет, выходим
	if len(swears) == 0 {
		return nil
	}

	// Сохраняем статистику
	if err = s.statisticService.SaveStatistic(ctx, statisticModel.SaveStatisticsReq{
		UserID:    user.ID,
		MessageID: message.ID,
		ChatID:    message.Chat.ID,
		Swears:    swears,
		Datetime:  message.Time(),
	}); err != nil {
		return err
	}

	// Если маты есть
	if len(swears) != 0 {

		// Ставим реакцию
		if err = s.tgBot.React(message.Chat, message, telebot.ReactionOptions{
			Reactions: []telebot.Reaction{{
				Type:        "emoji",
				Emoji:       "👾",
				CustomEmoji: "",
			}},
			Big: true,
		}); err != nil {
			return errors.InternalServer.Wrap(err)
		}
	}

	return nil
}
