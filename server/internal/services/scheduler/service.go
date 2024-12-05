package scheduler

import (
	"github.com/robfig/cron/v3"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("/server/internal/services/scheduler/service")

type Scheduler struct {
	cron *cron.Cron
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		cron: cron.New(),
	}
}
