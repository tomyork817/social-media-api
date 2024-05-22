package inmemory

import (
	"social-media-api/internal/models"
	"sync"
)

type CommentInMemory struct {
	data   map[int]*models.Comment
	lastID int
	lock   sync.Mutex
}
