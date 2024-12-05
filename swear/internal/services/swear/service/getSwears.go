package service

import (
	"context"

	"pkg/slices"
)

func (s *SwearService) GetSwears(ctx context.Context) (map[string]struct{}, error) {
	ctx, span := tracer.Start(ctx, "GetSwears")
	defer span.End()

	// Получаем все маты
	swearsDictionary, err := s.swearRepository.GetSwears(ctx)
	if err != nil {
		return nil, err
	}

	// Делаем мапу
	return slices.GetMapValueStruct(swearsDictionary, func(swear string) string { return swear }), nil
}
