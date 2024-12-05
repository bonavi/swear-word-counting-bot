package model

type SaveStatisticsReq struct {
	UserID    int64
	ChatID    int64
	MessageID int
	Swears    []string
}
