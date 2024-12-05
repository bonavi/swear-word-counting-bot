package endpoint

import (
	"context"

	"gopkg.in/telebot.v3"

	"pkg/errors"
	"pkg/log"
)

// Чекинг сообщения
func (e *endpoint) commandStart(c telebot.Context) error {
	ctx := context.Background()

	tracer.Start(ctx, "commandStart")

	log.Info(ctx, "/start")

	if err := c.Send("hi"); err != nil {
		log.Error(ctx, errors.InternalServer.Wrap(err))
	}

	return nil
}
