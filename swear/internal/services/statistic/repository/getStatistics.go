package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"pkg/ddlHelper"
	"swearBot/internal/ddl/statisticDDL"
	"swearBot/internal/services/statistic/model"
	repoModel "swearBot/internal/services/statistic/repository/model"
)

func (r *StatisticRepository) GetStatistics(ctx context.Context, req model.GetStatisticsReq) (statistics []model.Statistic, err error) {
	ctx, span := tracer.Start(ctx, "GetStatistics")
	defer span.End()

	q := sq.
		Select(
			statisticDDL.ColumnUserID,
			statisticDDL.ColumnSwear,
			ddlHelper.As(
				ddlHelper.Count(ddlHelper.SelectAll),
				"count",
			),
		).
		From(statisticDDL.Table).
		GroupBy(
			statisticDDL.ColumnUserID,
			statisticDDL.ColumnSwear,
		).
		OrderBy(ddlHelper.Desc("count"))

	q = q.Where(sq.Eq{statisticDDL.ColumnChatID: req.ChatID})
	q = q.Where(sq.Gt{statisticDDL.ColumnDatetime: req.DateFrom})
	q = q.Where(sq.Lt{statisticDDL.ColumnDatetime: req.DateTo})

	var repoStatistics []repoModel.Statistic

	// Получаем плоскую статистику
	if err = r.db.Select(ctx, &repoStatistics, q)
		err != nil {
		return nil, err
	}

	swearStatisticByUser := make(map[int][]model.SwearStatistic)

	// Делаем дерево
	for _, repoStatistic := range repoStatistics {
		userStatistics := swearStatisticByUser[repoStatistic.UserID]
		userStatistics = append(userStatistics, model.SwearStatistic{
			Swear: repoStatistic.Swear,
			Count: repoStatistic.Count,
		})
		swearStatisticByUser[repoStatistic.UserID] = userStatistics
	}

	statistics = make([]model.Statistic, 0, len(swearStatisticByUser))

	for userID, swearStatistics := range swearStatisticByUser {
		statistics = append(statistics, model.Statistic{
			UserID:          userID,
			SwearStatistics: swearStatistics,
		})
	}

	return statistics, nil
}
