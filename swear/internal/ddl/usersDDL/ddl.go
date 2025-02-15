package usersDDL

import "swearBot/internal/ddl"

const (
	Table          = ddl.SchemaSwearsCountingBot + "." + "users"
	TableWithAlias = Table + " " + alias
	alias          = "m"
)

const (
	ColumnID                = "id"
	ColumnFirstName         = "first_name"
	ColumnLastName          = "last_name"
	ColumnIsForum           = "is_forum"
	ColumnUsername          = "username"
	ColumnLanguageCode      = "language_code"
	ColumnIsBot             = "is_bot"
	ColumnIsPremium         = "is_premium"
	ColumnAddedToMenu       = "added_to_menu"
	ColumnUsernames         = "usernames"
	ColumnCustomEmojiStatus = "custom_emoji_status"
)

func WithPrefix(column string) string {
	return alias + "." + column
}
