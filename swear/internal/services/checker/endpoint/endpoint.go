package endpoint

import (
	"context"

	"go.opentelemetry.io/otel"
	"gopkg.in/telebot.v3"

	"swearBot/internal/services/checker/model"
	checkerService "swearBot/internal/services/checker/service"
)

var tracer = otel.Tracer("/server/internal/services/checker/endpoint")

type endpoint struct {
	service CheckerService
	tgBot   *telebot.Bot
}

var _ CheckerService = new(checkerService.CheckerService)

type CheckerService interface {
	CheckMessage(context.Context, model.CheckMessageReq) (model.CheckMessageRes, error)
}

func NewCheckerEndpoint(tgBot *telebot.Bot, service CheckerService) {

	e := endpoint{
		service: service,
		tgBot:   tgBot,
	}

	tgBot.Handle(telebot.OnText, e.checkMessage) // Обработка обычного текста

}
