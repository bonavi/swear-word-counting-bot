package endpoint

import (
	"context"

	"gopkg.in/telebot.v3"

	"pkg/errors"
	"swearBot/internal/services/checker/model"
)

// Чекинг сообщения
func (e *endpoint) checkMessage(c telebot.Context) error {

	// Формируем базовый контекст
	ctx := context.Background()

	// Трейсинг
	ctx, span := tracer.Start(ctx, "checkMessage")
	defer span.End()

	// Получаем текст сообщения
	user := c.Sender()
	message := c.Message()

	// Проверяем сообщение и получаем на него ответ
	reply, err := e.service.CheckMessage(ctx, model.CheckMessageReq{
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
	})
	if err != nil {
		return err
	}

	// Если ответ есть
	if reply != "" {

		// Отправляем ответ
		if err = c.Send(reply); err != nil {
			return errors.InternalServer.Wrap(err)
		}
	}

	return nil
}
