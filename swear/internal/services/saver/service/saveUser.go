package service

import (
	"context"

	saverModel "swearBot/internal/services/saver/model"
)

func (s *SaverService) SaveUser(ctx context.Context, req saverModel.SaveUserReq) error {
	return s.saverRepository.SaveUser(ctx, req)
}
