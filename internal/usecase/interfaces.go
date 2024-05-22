package usecase

import (
	"context"
	"social-media-api/internal/models"
)

type (
	PostRepo interface {
		Save(ctx context.Context, post models.Post) (*models.Post, error)
		GetAll(ctx context.Context, limit, offset int) ([]*models.Post, error)
		GetByID(ctx context.Context, id int) (*models.Post, error)
		Update(ctx context.Context, comment models.Post) (*models.Post, error)
	}

	CommentRepo interface {
		Save(ctx context.Context, post models.Comment) (*models.Comment, error)
		// GetAll(ctx context.Context, limit, offset int) ([]*models.Comment, error)
		GetByPostID(ctx context.Context, postID int, limit, offset int) ([]*models.Comment, error)
		GetByParentID(ctx context.Context, parentID int, limit, offset int) ([]*models.Comment, error)
		GetByID(ctx context.Context, id int) (*models.Comment, error)
	}

	Post interface {
		Create(ctx context.Context, post models.Post) (*models.Post, error)
		GetAll(ctx context.Context) ([]*models.Post, error)
	}

	Comment interface {
	}
)
