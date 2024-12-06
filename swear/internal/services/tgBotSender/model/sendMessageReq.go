package model

import "gopkg.in/telebot.v3"

type SendMessageReq struct {
	Chat    *telebot.Chat // Чат
	Message string        // Сообщение
}
