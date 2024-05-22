package models

type PostFilter struct {
	ID int
}

type CommentFilter struct {
	ID       int
	ParentID int
	PostID   int
}
