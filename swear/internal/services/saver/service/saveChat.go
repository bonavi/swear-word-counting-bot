package service

import (
	"context"

	saverModel "swearBot/internal/services/saver/model"
)

func (s *SaverService) SaveChat(ctx context.Context, req saverModel.SaveChatReq) error {
	return s.saverRepository.SaveChat(ctx, req)
}
