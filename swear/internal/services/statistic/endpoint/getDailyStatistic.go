package endpoint

import (
	"context"
	"strings"

	"gopkg.in/telebot.v3"

	"pkg/datetime"
	"pkg/errors"
	"swearBot/internal/services/statistic/model"
)

// Чекинг сообщения
func (e *endpoint) getDailyStatistic(c telebot.Context) (err error) {

	// Формируем базовый контекст
	ctx := context.Background()

	// Трейсинг
	ctx, span := tracer.Start(ctx, "getDailyStatistic")
	defer span.End()

	// Получаем текст сообщения
	message := c.Message()

	var day datetime.Date

	// Получаем день
	dayString := message.Payload

	// Парсим день
	switch strings.ToLower(dayString) {
	case "today", "сегодня", "":
		day = datetime.Today()
	case "yesterday", "вчера":
		day = datetime.Today().AddDate(0, 0, -1)
	default:
		day, err = datetime.Parse(dayString)
		if err != nil {
			if err := c.Send("Неверный формат даты, ожидается YYYY-MM-DD"); err != nil {
				return err
			}
			return err
		}
	}

	// Проверяем сообщение и получаем на него ответ
	reply, err := e.service.GetDailyStatisticReply(ctx, model.GetStatisticsReq{
		ChatID:   message.Chat.ID,
		DateFrom: day,
		DateTo:   day.AddDate(0, 0, 1),
	})
	if err != nil {
		return err
	}

	// Отправляем ответ
	if err = c.Send(reply); err != nil {
		return errors.InternalServer.Wrap(err)
	}

	return nil
}
