package usecase

import (
	"context"
	"social-media-api/internal/domain"
)

type (
	PostRepo interface {
		Save(ctx context.Context, post domain.Post) (*domain.Post, error)
		GetAll(ctx context.Context) ([]*domain.Post, error)
	}

	Post interface {
		Create(ctx context.Context, post domain.Post) (*domain.Post, error)
		GetAll(ctx context.Context) ([]*domain.Post, error)
	}
)
