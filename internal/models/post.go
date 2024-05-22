package models

const (
	PostBodyMaxLength = 3000
)

type Post struct {
	ID     int    `db:"id"`
	UserID int    `db:"user_id"`
	Body   string `db:"body"`
	IsOpen bool   `db:"is_open"`
}
