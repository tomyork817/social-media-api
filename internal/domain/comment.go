package domain

type Comment struct {
	ID       int
	UserID   int
	PostID   int
	ParentID int
	Body     string
}
