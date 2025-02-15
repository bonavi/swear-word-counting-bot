package service

import (
	"context"

	"go.opentelemetry.io/otel"
	"gopkg.in/telebot.v3"

	saverModel "swearBot/internal/services/saver/model"
	saverService "swearBot/internal/services/saver/service"
	statisticModel "swearBot/internal/services/statistic/model"
	statisticService "swearBot/internal/services/statistic/service"
	swearService "swearBot/internal/services/swear/service"
)

var tracer = otel.Tracer("/server/internal/services/checker/service")

type CheckerService struct {
	statisticService StatisticService
	swearService     SwearService
	saverService     SaverService
	tgBot            *telebot.Bot
}

var _ StatisticService = new(statisticService.StatisticService)

type StatisticService interface {
	SaveStatistic(context.Context, statisticModel.SaveStatisticsReq) error
}

var _ SwearService = new(swearService.SwearService)

type SwearService interface {
	GetSwears(context.Context) (map[string]struct{}, error)
}

var _ SaverService = new(saverService.SaverService)

type SaverService interface {
	SaveUser(context.Context, saverModel.SaveUserReq) error
	SaveChat(context.Context, saverModel.SaveChatReq) error
	SaveMessage(context.Context, saverModel.SaveMessageReq) error
}

func NewCheckerService(
	checkerService StatisticService,
	swearService SwearService,
	saverService SaverService,
	tgBot *telebot.Bot,
) {

	e := &CheckerService{
		statisticService: checkerService,
		swearService:     swearService,
		saverService:     saverService,
		tgBot:            tgBot,
	}

	tgBot.Handle(telebot.OnText, e.CheckMessage) // Обработка обычного текста
}
