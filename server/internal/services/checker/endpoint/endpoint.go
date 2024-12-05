package endpoint

import (
	"context"

	"go.opentelemetry.io/otel"
	"gopkg.in/telebot.v3"

	"server/internal/services/checker/model"
	checkerService "server/internal/services/checker/service"
)

var tracer = otel.Tracer("/server/internal/services/checker/endpoint")

type endpoint struct {
	service CheckerService
}

var _ CheckerService = new(checkerService.CheckerService)

type CheckerService interface {
	CheckMessage(context.Context, model.CheckMessageReq) error
}

func NewTgBotEndpoint(tgBot *telebot.Bot, service CheckerService) {

	e := endpoint{
		service: service,
	}

	tgBot.Handle("/start", e.commandStart)

	tgBot.Handle(telebot.OnText, e.checkMessage) // Обработка обычного текста

}
