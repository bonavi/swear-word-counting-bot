package service

import (
	"context"
	"strings"

	"pkg/slices"
	"swearBot/internal/services/swear/model"
)

func (s *SwearService) AddSwears(ctx context.Context, req model.AddSwearsReq) (int, error) {
	ctx, span := tracer.Start(ctx, "AddSwears")
	defer span.End()

	// Приводим каждое слово к нижнему регистру
	for i, swear := range req.Swears {
		req.Swears[i] = strings.ToLower(swear)
	}

	// Получаем маты из базы данных
	swearsMap, err := s.GetSwears(ctx)
	if err != nil {
		return 0, err
	}

	var newSwears []string

	// Проходимся по каждому слову
	for _, swear := range req.Swears {

		// Если слова у нас нет
		if _, ok := swearsMap[swear]; !ok {

			// Добавляем в список новых слов
			newSwears = append(newSwears, swear)
		}
	}

	req.Swears = slices.Unique(newSwears)

	if len(req.Swears) == 0 {
		return 0, nil
	}

	// Добавляем маты
	if err = s.swearRepository.AddSwears(ctx, req); err != nil {
		return 0, err
	}

	return len(req.Swears), nil
}
