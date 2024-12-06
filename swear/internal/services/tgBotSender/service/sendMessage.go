package service

import (
	"context"

	"pkg/errors"
	"pkg/log"

	"swearBot/internal/services/tgBotSender/model"
)

// SendMessage отправляет сообщение пользователю в телеграм
func (s *TgBotSenderService) SendMessage(ctx context.Context, req model.SendMessageReq) error {

	ctx, span := tracer.Start(ctx, "SendMessage")
	defer span.End()

	if !s.isOn {
		log.Warning(ctx, "Вызвана функция SendMessage. Отсылка сообщений выключена")
		return nil
	}

	if _, err := s.Bot.Send(req.Chat, req.Message); err != nil {
		return errors.InternalServer.Wrap(err)
	}

	return nil
}