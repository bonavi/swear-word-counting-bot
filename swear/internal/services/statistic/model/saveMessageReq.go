package model

import "time"

type SaveMessageReq struct {
	UserID    int64
	ChatID    int64
	MessageID int
	Swears    []string
	Datetime  time.Time
}
