package scheduler

func (s *Scheduler) Start() error {

	s.cron.Start()

	return nil
}
