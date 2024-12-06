package model

import "pkg/datetime"

type GetStatisticsRes struct {
	Date       datetime.Date
	Statistics []Statistic
}
