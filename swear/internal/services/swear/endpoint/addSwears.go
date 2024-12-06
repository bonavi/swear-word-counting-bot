package endpoint

import (
	"context"
	"strings"

	"gopkg.in/telebot.v3"

	"pkg/log"
	"swearBot/internal/services/swear/model"
	tgBotSenderModel "swearBot/internal/services/tgBotSender/model"
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
		Chat:     c.Chat(),
	}); err != nil {
		log.Error(ctx, err)
	}

	// Отправляем сообщение о добавлении слов
	if err := e.tgBotSenderService.SendMessage(ctx, tgBotSenderModel.SendMessageReq{
		Chat:    c.Chat(),
		Message: "✅ Слова добавлены",
	}); err != nil {
		log.Error(ctx, err)
	}

	return nil
}
