package model

import "time"

type Message struct {
	ID     int
	ChatID int64
	Text   string
	Time   time.Time
}
