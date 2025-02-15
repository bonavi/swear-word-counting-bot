package service

import (
	"context"

	"go.opentelemetry.io/otel"

	saverModel "swearBot/internal/services/saver/model"
	saverRepository "swearBot/internal/services/saver/repository"
)

var tracer = otel.Tracer("/server/internal/services/saver/service") //nolint:unused

type SaverService struct {
	saverRepository SaverRepository
}

var _ SaverRepository = new(saverRepository.SaverRepository)

type SaverRepository interface {
	SaveUser(context.Context, saverModel.SaveUserReq) error
	SaveChat(context.Context, saverModel.SaveChatReq) error
	SaveMessage(context.Context, saverModel.SaveMessageReq) error
}

func NewSaverService(
	saverRepository SaverRepository,
) *SaverService {
	return &SaverService{
		saverRepository: saverRepository,
	}
}
