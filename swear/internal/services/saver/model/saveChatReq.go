package model

type SaveChatReq struct {
	ID                       int64
	Type                     string
	Title                    string
	FirstName                string
	LastName                 string
	Username                 string
	Bio                      string
	Description              string
	InviteLink               string
	SlowMode                 int
	StickerSet               string
	CanSetStickerSet         bool
	CustomEmojiSetName       string
	LinkedChatID             int64
	Private                  bool
	Protected                bool
	NoVoiceAndVideo          bool
	HasHiddenMembers         bool
	AggressiveAntiSpam       bool
	CustomEmojiID            string
	EmojiExpirationUnixtime  int64
	BackgroundEmojiID        string
	AccentColorID            int
	ProfileAccentColorID     int
	ProfileBackgroundEmojiID string
	HasVisibleHistory        bool
	UnrestrictBoosts         int
}
