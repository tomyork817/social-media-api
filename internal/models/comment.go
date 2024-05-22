package models

const (
	CommentBodyMaxLength = 2000
)

type Comment struct {
	ID       int    `db:"id"`
	UserID   int    `db:"user_id"`
	PostID   int    `db:"post_id"`
	ParentID int    `db:"parent_id"`
	Body     string `db:"body"`
}
