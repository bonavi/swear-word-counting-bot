package model

type Statistic struct {
	UserID          int              `json:"userID"`
	SwearStatistics []SwearStatistic `json:"swearStatistics"`
}

type SwearStatistic struct {
	Swear string `json:"swear"`
	Count int    `json:"count"`
}
