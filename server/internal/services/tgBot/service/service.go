package service

import (
	"context"

	"go.opentelemetry.io/otel"
	"gopkg.in/telebot.v3"

	"pkg/log"
)

var tracer = otel.Tracer("/server/internal/services/tgBot/service")

type TgBotService struct {
	Bot *telebot.Bot

	isOn bool
}

func NewTgBotService(
	tgBot *telebot.Bot,
	isOn bool,
) *TgBotService {

	if !isOn {
		log.Warning(context.Background(), "Telegram bot is off", log.SkipThisCallOption())
		return &TgBotService{
			Bot:  nil,
			isOn: isOn,
		}
	}

	return &TgBotService{
		Bot:  tgBot,
		isOn: isOn,
	}
}
