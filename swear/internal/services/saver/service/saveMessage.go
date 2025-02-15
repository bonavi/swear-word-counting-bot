package service

import (
	"context"

	saverModel "swearBot/internal/services/saver/model"
)

func (s *SaverService) SaveMessage(ctx context.Context, req saverModel.SaveMessageReq) error {
	return s.saverRepository.SaveMessage(ctx, req)
}
