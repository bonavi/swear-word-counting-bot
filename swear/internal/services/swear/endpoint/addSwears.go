package endpoint

import (
	"context"
	"fmt"
	"strings"

	"gopkg.in/telebot.v3"

	"pkg/log"
	"swearBot/internal/services/swear/model"
)

// Чекинг сообщения
func (e *endpoint) addSwears(c telebot.Context) error {

	// Формируем базовый контекст
	ctx := context.Background()

	log.Info(ctx, commandAdd)

	// Трейсинг
	ctx, span := tracer.Start(ctx, "addSwears")
	defer span.End()

	// Получаем текст сообщения
	message := c.Message()

	swearsLine := strings.ReplaceAll(message.Text, commandAdd, "")

	user := c.Sender()

	// Получаем слова
	swears := strings.Fields(swearsLine)

	// Проверяем сообщение
	count, err := e.swearService.AddSwears(ctx, model.AddSwearsReq{
		Swears:   swears,
		UserID:   user.ID,
		Datetime: message.Time(),
		Chat:     c.Chat(),
	})
	if err != nil {
		log.Error(ctx, err)
	}

	// Отправляем сообщение о добавлении слов
	if err = c.Send(
		fmt.Sprintf("✅ %d новых слов добавлено", count),
	); err != nil {
		log.Error(ctx, err)
	}

	return nil
}
