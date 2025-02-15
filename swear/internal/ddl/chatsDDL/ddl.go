package chatsDDL

import "swearBot/internal/ddl"

const (
	Table          = ddl.SchemaSwearsCountingBot + "." + "chats"
	TableWithAlias = Table + " " + alias
	alias          = "c"
)

const (
	ColumnID                       = "id"
	ColumnType                     = "type"
	ColumnTitle                    = "title"
	ColumnFirstName                = "first_name"
	ColumnLastName                 = "last_name"
	ColumnUsername                 = "username"
	ColumnBio                      = "bio"
	ColumnDescription              = "description"
	ColumnInviteLink               = "invite_link"
	ColumnSlowMode                 = "slow_mode"
	ColumnStickerSet               = "sticker_set"
	ColumnCanSetStickerSet         = "can_set_sticker_set"
	ColumnCustomEmojiSetName       = "custom_emoji_set_name"
	ColumnLinkedChatId             = "linked_chat_id"
	ColumnPrivate                  = "private"
	ColumnProtected                = "protected"
	ColumnNoVoiceAndVideo          = "no_voice_and_video"
	ColumnHasHiddenMembers         = "has_hidden_members"
	ColumnAggressiveAntiSpam       = "aggressive_anti_spam"
	ColumnCustomEmojiId            = "custom_emoji_id"
	ColumnEmojiExpirationUnixtime  = "emoji_expiration_unixtime"
	ColumnBackgroundEmojiId        = "background_emoji_id"
	ColumnAccentColorId            = "accent_color_id"
	ColumnProfileAccentColorId     = "profile_accent_color_id"
	ColumnProfileBackgroundEmojiId = "profile_background_emoji_id"
	ColumnHasVisibleHistory        = "has_visible_history"
	ColumnUnrestrictBoosts         = "unrestrict_boosts"
	ColumnCountUsers               = "count_users"
)

func WithPrefix(column string) string {
	return alias + "." + column
}
