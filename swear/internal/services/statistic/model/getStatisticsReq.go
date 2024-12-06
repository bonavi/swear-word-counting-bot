package model

import "pkg/datetime"

type GetStatisticsReq struct {
	ChatID int64
	DateFrom   datetime.Date
	DateTo   datetime.Date
}
