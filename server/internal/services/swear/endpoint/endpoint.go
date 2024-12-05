package endpoint

import (
	"context"

	"go.opentelemetry.io/otel"
	"gopkg.in/telebot.v3"

	"server/internal/services/swear/model"
	swearService "server/internal/services/swear/service"
)

var tracer = otel.Tracer("/server/internal/services/swear/endpoint")

type endpoint struct {
	swearService SwearService
}

var _ SwearService = new(swearService.SwearService)

type SwearService interface {
	AddSwears(context.Context, model.AddSwearsReq) error
}

func NewSwearEndpoint(tgBot *telebot.Bot, service SwearService) {

	e := endpoint{
		swearService: service,
	}

	tgBot.Handle("/addSwears", e.addSwears)
}
