package service

import (
	"context"

	"go.opentelemetry.io/otel"
	"gopkg.in/telebot.v3"

	"pkg/log"
)

var tracer = otel.Tracer("/server/internal/services/tgBot/service")

type TgBotSenderService struct {
	Bot *telebot.Bot

	isOn bool
}

func NewTgBotSenderService(
	tgBot *telebot.Bot,
	isOn bool,
) *TgBotSenderService {

	if !isOn {
		log.Warning(context.Background(), "Telegram bot is off", log.SkipThisCallOption())
		return &TgBotSenderService{
			Bot:  nil,
			isOn: isOn,
		}
	}

	return &TgBotSenderService{
		Bot:  tgBot,
		isOn: isOn,
	}
}
