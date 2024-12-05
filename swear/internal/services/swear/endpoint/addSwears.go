package endpoint

import (
	"context"
	"strings"

	"gopkg.in/telebot.v3"

	"pkg/log"
	"swearBot/internal/services/swear/model"
)

// Чекинг сообщения
func (e *endpoint) addSwears(c telebot.Context) error {

	// Формируем базовый контекст
	ctx := context.Background()

	log.Info(ctx, "/addSwears")

	// Трейсинг
	ctx, span := tracer.Start(ctx, "addSwears")
	defer span.End()

	// Получаем текст сообщения
	message := c.Message()
	swearsLine := message.Payload
	user := c.Sender()

	// Получаем слова
	swears := strings.Fields(swearsLine)

	// Приводим каждое слово к нижнему регистру
	for i, swear := range swears {
		swears[i] = strings.ToLower(swear)
	}

	// Проверяем сообщение
	if err := e.swearService.AddSwears(ctx, model.AddSwearsReq{
		Swears:   swears,
		UserID:   user.ID,
		Datetime: message.Time(),
	}); err != nil {
		log.Error(ctx, err)
	}

	return nil
}
