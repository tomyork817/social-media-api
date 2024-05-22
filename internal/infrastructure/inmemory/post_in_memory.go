package inmemory

import (
	"cmp"
	"context"
	"slices"
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

func (r *PostInMemory) GetAll(ctx context.Context, limit, offset int) ([]*models.Post, error) {
	posts := make([]*models.Post, 0, len(r.data))
	r.lock.Lock()
	for _, post := range r.data {
		posts = append(posts, post)
	}
	r.lock.Unlock()
	slices.SortFunc(posts, func(a, b *models.Post) int {
		return cmp.Compare(a.ID, b.ID)
	})

	start := offset
	end := start + limit
	if start > len(posts) {
		start = len(posts)
		end = len(posts)
	}
	if end > len(posts) {
		end = len(posts)
	}

	return posts[start:end], nil
}

func (r *PostInMemory) GetByID(ctx context.Context, id int) (*models.Post, error) {
	r.lock.Lock()
	post := r.data[id]
	r.lock.Unlock()
	if post == nil {
		return nil, models.ErrNotFound
	}
	return post, nil
}

func (r *PostInMemory) UpdateIsOpenById(ctx context.Context, id int, isOpen bool) (*models.Post, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	_, ok := r.data[id]
	if !ok {
		return nil, models.ErrNotFound
	}
	r.data[id].IsOpen = isOpen

	return r.data[id], nil
}

func (r *PostInMemory) GetByUserID(ctx context.Context, userID int, limit, offset int) ([]*models.Post, error) {
	posts := make([]*models.Post, 0)
	r.lock.Lock()
	for _, post := range r.data {
		if post.UserID == userID {
			posts = append(posts, post)
		}
	}
	r.lock.Unlock()

	slices.SortFunc(posts, func(a, b *models.Post) int {
		return cmp.Compare(a.ID, b.ID)
	})

	start := offset
	end := start + limit
	if start > len(posts) {
		start = len(posts)
		end = len(posts)
	}
	if end > len(posts) {
		end = len(posts)
	}

	return posts[start:end], nil
}
