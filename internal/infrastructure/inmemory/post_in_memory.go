package inmemory

import (
	"context"
	"social-media-api/internal/models"
	"sync"
)

type PostInMemory struct {
	data   map[int]*models.Post
	lastID int
	lock   sync.Mutex
}

func NewPostInMemory() *PostInMemory {
	return &PostInMemory{
		data: make(map[int]*models.Post),
	}
}

func (r *PostInMemory) Save(ctx context.Context, post models.Post) (*models.Post, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.lastID++
	post.ID = r.lastID
	r.data[r.lastID] = &post

	return r.data[r.lastID], nil
}

func (r *PostInMemory) GetAll(ctx context.Context) ([]*models.Post, error) {
	posts := make([]*models.Post, 0, len(r.data))
	for _, post := range r.data {
		posts = append(posts, post)
	}

	return posts, nil
}
