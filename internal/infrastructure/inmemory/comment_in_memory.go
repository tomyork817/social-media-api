package inmemory

import (
	"cmp"
	"context"
	"slices"
	"social-media-api/internal/models"
	"sync"
)

type CommentInMemory struct {
	data   map[int]*models.Comment
	lastID int
	lock   sync.Mutex
}

func NewCommentInMemory() *CommentInMemory {
	return &CommentInMemory{
		data: make(map[int]*models.Comment),
	}
}

func (r *CommentInMemory) Save(ctx context.Context, comment models.Comment) (*models.Comment, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.lastID++
	comment.ID = r.lastID
	r.data[r.lastID] = &comment

	return r.data[r.lastID], nil
}

func (r *CommentInMemory) GetByPostID(ctx context.Context, postID int, limit, offset int) ([]*models.Comment, error) {
	comments := make([]*models.Comment, 0)
	r.lock.Lock()
	for _, comment := range r.data {
		if comment.PostID == postID && comment.ParentID == 0 {
			comments = append(comments, comment)
		}
	}
	r.lock.Unlock()

	slices.SortFunc(comments, func(a, b *models.Comment) int {
		return cmp.Compare(a.ID, b.ID)
	})

	start := offset
	end := start + limit
	if start > len(comments) {
		start = len(comments)
		end = len(comments)
	}
	if end > len(comments) {
		end = len(comments)
	}

	return comments[start:end], nil
}

func (r *CommentInMemory) GetByParentID(ctx context.Context, parentID int, limit, offset int) ([]*models.Comment, error) {
	comments := make([]*models.Comment, 0)
	r.lock.Lock()
	for _, comment := range r.data {
		if comment.ParentID == parentID {
			comments = append(comments, comment)
		}
	}
	r.lock.Unlock()

	slices.SortFunc(comments, func(a, b *models.Comment) int {
		return cmp.Compare(a.ID, b.ID)
	})

	start := offset
	end := start + limit
	if start > len(comments) {
		start = len(comments)
		end = len(comments)
	}
	if end > len(comments) {
		end = len(comments)
	}

	return comments[start:end], nil
}

func (r *CommentInMemory) GetByID(ctx context.Context, id int) (*models.Comment, error) {
	r.lock.Lock()
	comment := r.data[id]
	r.lock.Unlock()
	if comment == nil {
		return nil, models.ErrNotFound
	}
	return comment, nil
}

func (r *CommentInMemory) GetAll(ctx context.Context, limit, offset int) ([]*models.Comment, error) {
	comments := make([]*models.Comment, 0)
	r.lock.Lock()
	for _, comment := range r.data {
		comments = append(comments, comment)
	}
	r.lock.Unlock()

	slices.SortFunc(comments, func(a, b *models.Comment) int {
		return cmp.Compare(a.ID, b.ID)
	})

	start := offset
	end := start + limit
	if start > len(comments) {
		start = len(comments)
		end = len(comments)
	}
	if end > len(comments) {
		end = len(comments)
	}

	return comments[start:end], nil
}
