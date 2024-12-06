package model

type Statistic struct {
	UserID int    `db:"user_id"`
	Swear  string `db:"swear"`
	Count  int    `db:"count"`
}
