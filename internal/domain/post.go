package domain

type Post struct {
	ID     int
	UserID int
	Body   string
	IsOpen bool
}
