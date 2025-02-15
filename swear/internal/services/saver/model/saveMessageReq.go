package model

import "time"

type SaveMessageReq struct {
	ID                 int
	ChatID             int64
	UserID             int64
	ThreadID           int
	DateTime           time.Time
	OriginalMessageID  int
	OriginalSignature  string
	OriginalSenderName string
	OriginalUnixtime   int
	AutomaticForward   bool
	LastEdit           int64
	TopicMessage       bool
	Protected          bool
	AlbumID            string
	Signature          string
	Text               string
	Payload            string
	Caption            string
	NewGroupTitle      string
	GroupPhotoDeleted  bool
	GroupCreated       bool
	SuperGroupCreated  bool
	ChannelCreated     bool
	MigrateTo          int64
	MigrateFrom        int64
	ConnectedWebsite   string
	SenderBoostCount   int
	HasMediaSpoiler    bool
}
