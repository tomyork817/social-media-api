package models

const (
	CommentBodyMaxLength = 2000
)

type Comment struct {
	ID       int
	UserID   int
	PostID   int
	ParentID int
	Body     string
}
