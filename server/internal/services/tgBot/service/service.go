package service

import (
	"context"

	"go.opentelemetry.io/otel"
	"gopkg.in/telebot.v3"

	"pkg/errors"
	"pkg/log"
)

var tracer = otel.Tracer("/server/internal/services/tgBot/service")

type TgBotService struct {
	Bot *telebot.Bot

	isOn bool
}

func NewTgBotService(
	token string,
	isOn bool,
) (*TgBotService, error) {

	if !isOn {
		log.Warning(context.Background(), "Telegram bot is off", log.SkipThisCallOption())
		return &TgBotService{
			Bot:  nil,
			isOn: isOn,
		}, nil
	}

	bot, err := telebot.NewBot(telebot.Settings{
		URL:         "",
		Token:       token,
		Updates:     0,
		Poller:      nil,
		Synchronous: false,
		Verbose:     false,
		ParseMode:   telebot.ModeHTML,
		OnError:     nil,
		Client:      nil,
		Offline:     false,
	})
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	return &TgBotService{
		Bot:  bot,
		isOn: isOn,
	}, nil
}
