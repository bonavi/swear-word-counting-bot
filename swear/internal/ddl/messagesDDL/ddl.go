package messagesDDL

import "swearBot/internal/ddl"

const (
	Table          = ddl.SchemaSwearsCountingBot + "." + "messages"
	TableWithAlias = Table + " " + alias
	alias          = "m"
)

const (
	ColumnID                 = "id"
	ColumnChatID             = "chat_id"
	ColumnUserID             = "user_id"
	ColumnThreadID           = "thread_id"
	ColumnDateTime           = "date_time"
	ColumnOriginalMessageID  = "original_message_id"
	ColumnOriginalSignature  = "original_signature"
	ColumnOriginalSenderName = "original_sender_name"
	ColumnOriginalUnixtime   = "original_unixtime"
	ColumnAutomaticForward   = "automatic_forward"
	ColumnLastEdit           = "last_edit"
	ColumnTopicMessage       = "topic_message"
	ColumnProtected          = "protected"
	ColumnAlbumID            = "album_id"
	ColumnSignature          = "signature"
	ColumnText               = "text"
	ColumnPayload            = "payload"
	ColumnCaption            = "caption"
	ColumnNewGroupTitle      = "new_group_title"
	ColumnGroupPhotoDeleted  = "group_photo_deleted"
	ColumnGroupCreated       = "group_created"
	ColumnSuperGroupCreated  = "super_group_created"
	ColumnChannelCreated     = "channel_created"
	ColumnMigrateTo          = "migrate_to"
	ColumnMigrateFrom        = "migrate_from"
	ColumnConnectedWebsite   = "connected_website"
	ColumnSenderBoostCount   = "sender_boost_count"
	ColumnHasMediaSpoiler    = "has_media_spoiler"
)

func WithPrefix(column string) string {
	return alias + "." + column
}
