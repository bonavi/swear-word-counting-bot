package endpoint

import (
	"context"

	"go.opentelemetry.io/otel"
	"gopkg.in/telebot.v3"

	"swearBot/internal/services/statistic/model"
	statisticService "swearBot/internal/services/statistic/service"
)

var tracer = otel.Tracer("/server/internal/services/statistic/endpoint")

type endpoint struct {
	service StatisticService
}

var _ StatisticService = new(statisticService.StatisticService)

type StatisticService interface {
	GetDailyStatisticReply(context.Context, model.GetStatisticsReq) (string, error)
}

func NewStatisticEndpoint(tgBot *telebot.Bot, service StatisticService) {

	e := endpoint{
		service: service,
	}

	tgBot.Handle(commandGetStat, e.getDailyStatistic) // Получение дневной статистики

}

const commandGetStat = "/dailyStat"
