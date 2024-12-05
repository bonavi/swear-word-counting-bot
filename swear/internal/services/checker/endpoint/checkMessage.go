package endpoint

import (
	"context"

	"gopkg.in/telebot.v3"

	"pkg/log"
	"swearBot/internal/services/checker/model"
)

// Чекинг сообщения
func (e *endpoint) checkMessage(c telebot.Context) error {

	// Формируем базовый контекст
	ctx := context.Background()

	log.Info(ctx, "Хендлим сообщение")

	// Трейсинг
	ctx, span := tracer.Start(ctx, "checkMessage")
	defer span.End()

	// Получаем текст сообщения
	user := c.Sender()
	message := c.Message()

	// Проверяем сообщение
	if err := e.service.CheckMessage(ctx, model.CheckMessageReq{
		User: model.User{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
		Message: model.Message{
			Text:     message.Text,
			Datetime: message.Time(),
			ID:       message.ID,
			ChatID:   message.Chat.ID,
		},
	}); err != nil {
		log.Error(ctx, err)
	}

	return nil
}
