package model

type SaveUserReq struct {
	ID                int64
	FirstName         string
	LastName          string
	IsForum           bool
	Username          string
	LanguageCode      string
	IsBot             bool
	IsPremium         bool
	AddedToMenu       bool
	Usernames         []string
	CustomEmojiStatus string
}
