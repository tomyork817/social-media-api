package models

const (
	PostBodyMaxLength = 3000
)

type Post struct {
	ID     int
	UserID int
	Body   string
	IsOpen bool
}
