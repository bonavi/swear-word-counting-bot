package statisticDDL

import "swearBot/internal/ddl"

const (
	Table          = ddl.SchemaSwearsCountingBot + "." + "statistics"
	TableWithAlias = Table + " " + alias
	alias          = "st"
)

const (
	ColumnID        = "id"
	ColumnMessageID = "message_id"
	ColumnChatID    = "chat_id"
	ColumnUserID    = "user_id"
	ColumnSwear     = "swear"
	ColumnDatetime  = "datetime"
)

func WithPrefix(column string) string {
	return alias + "." + column
}
