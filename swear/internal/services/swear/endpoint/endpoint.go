package endpoint

import (
	"context"

	"go.opentelemetry.io/otel"
	"gopkg.in/telebot.v3"

	"swearBot/internal/services/swear/model"
	swearService "swearBot/internal/services/swear/service"
)

var tracer = otel.Tracer("/server/internal/services/swear/endpoint")

type endpoint struct {
	swearService SwearService
}

var _ SwearService = new(swearService.SwearService)

type SwearService interface {
	AddSwears(context.Context, model.AddSwearsReq) (int, error)
}

func NewSwearEndpoint(
	tgBot *telebot.Bot,
	service SwearService,
) {

	e := endpoint{
		swearService: service,
	}

	tgBot.Handle(commandAdd, e.addSwears)
}

const commandAdd = "/add"
