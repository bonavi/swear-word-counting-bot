package model

import "time"

type AddSwearsReq struct {
	Swears   []string
	UserID   int64
	Datetime time.Time
}
