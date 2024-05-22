package models

type PostFilter struct {
	UserID int
}

type CommentFilter struct {
	ParentID int
	PostID   int
}
