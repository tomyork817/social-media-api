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
		GetByUserID(ctx context.Context, userID int, limit, offset int) ([]*models.Post, error)
		UpdateIsOpenById(ctx context.Context, id int, isOpen bool) (*models.Post, error)
	}

	CommentRepo interface {
		Save(ctx context.Context, comment models.Comment) (*models.Comment, error)
		GetAll(ctx context.Context, limit, offset int) ([]*models.Comment, error)
		GetByPostID(ctx context.Context, postID int, limit, offset int) ([]*models.Comment, error)
		GetByParentID(ctx context.Context, parentID int, limit, offset int) ([]*models.Comment, error)
		GetByID(ctx context.Context, id int) (*models.Comment, error)
	}

	Post interface {
		Create(ctx context.Context, post models.Post) (*models.Post, error)
		GetById(ctx context.Context, id int) (*models.Post, error)
		GetMultiple(ctx context.Context, filter models.PostFilter, limit, offset int) ([]*models.Post, error)
		DisableComments(ctx context.Context, id int) (*models.Post, error)
		EnableComments(ctx context.Context, id int) (*models.Post, error)
	}

	Comment interface {
		Create(ctx context.Context, comment models.Comment) (*models.Comment, error)
		GetById(ctx context.Context, id int) (*models.Comment, error)
		GetMultiple(ctx context.Context, filter models.CommentFilter, limit, offset int) ([]*models.Comment, error)
	}
)
