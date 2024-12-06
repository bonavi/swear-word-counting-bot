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
	res, err := e.service.CheckMessage(ctx, model.CheckMessageReq{
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

	if res.SwearsCount != 0 {
		if err = e.tgBot.React(message.Chat, message, telebot.ReactionOptions{
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

	// Если ответ есть
	if res.Reply != "" {

		// Отправляем ответ
		if err = c.Send(res.Reply); err != nil {
			return errors.InternalServer.Wrap(err)
		}
	}

	return nil
}
