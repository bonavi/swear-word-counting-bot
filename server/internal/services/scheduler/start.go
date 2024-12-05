package scheduler

import (
	"context"
	"time"

	"pkg/errors"
)

func (s *Scheduler) Start() error {

	// Обновление валют
	_, err := s.cron.AddFunc("@mountly", func() { // Every month at first day of month 00:00 UTC
		_, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

	})
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}

	s.cron.Start()

	return nil
}
