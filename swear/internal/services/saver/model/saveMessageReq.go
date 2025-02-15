package model

type SaveMessageReq struct {
	ID                 int
	ThreadID           int
	Unixtime           int64
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
