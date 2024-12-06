package endpoint

import (
	"context"

	"go.opentelemetry.io/otel"
	"gopkg.in/telebot.v3"

	"swearBot/internal/services/swear/model"
	swearService "swearBot/internal/services/swear/service"
	tgBotSenderModel "swearBot/internal/services/tgBotSender/model"
	tgBotService "swearBot/internal/services/tgBotSender/service"
)

var tracer = otel.Tracer("/server/internal/services/swear/endpoint")

type endpoint struct {
	swearService       SwearService
	tgBotSenderService TgBotSenderService
}

var _ TgBotSenderService = new(tgBotService.TgBotSenderService)

type TgBotSenderService interface {
	SendMessage(ctx context.Context, req tgBotSenderModel.SendMessageReq) error
}

var _ SwearService = new(swearService.SwearService)

type SwearService interface {
	AddSwears(context.Context, model.AddSwearsReq) (int, error)
}

func NewSwearEndpoint(
	tgBot *telebot.Bot,
	tgBotSenderService TgBotSenderService,
	service SwearService,
) {

	e := endpoint{
		swearService:       service,
		tgBotSenderService: tgBotSenderService,
	}

	tgBot.Handle("/add", e.addSwears)
}
