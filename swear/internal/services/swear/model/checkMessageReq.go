package model

import (
	"time"

	"gopkg.in/telebot.v3"
)

type AddSwearsReq struct {
	Chat     *telebot.Chat
	Swears   []string
	UserID   int64
	Datetime time.Time
}
