package service

import (
	"context"

	"swearBot/internal/services/swear/model"
)

func (s *SwearService) AddSwears(ctx context.Context, req model.AddSwearsReq) error {
	ctx, span := tracer.Start(ctx, "AddSwears")
	defer span.End()

	// Добавляем маты
	return s.swearRepository.AddSwears(ctx, req)
}
