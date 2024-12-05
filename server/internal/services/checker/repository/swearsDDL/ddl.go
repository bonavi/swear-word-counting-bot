package swearsDDL

import "server/internal/ddl"

const (
	Table          = ddl.SchemaSwearsCountingBot + "." + "swears"
	TableWithAlias = Table + " " + alias
	alias          = "sw"
)

const (
	ColumnText = "text"
)

func WithPrefix(column string) string {
	return alias + "." + column
}
