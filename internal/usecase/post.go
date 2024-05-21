package usecase

import (
	"context"
	"social-media-api/internal/domain"
)

type PostUseCase struct {
	postRepo PostRepo
}

func NewPostUseCase(postRepo PostRepo) *PostUseCase {
	return &PostUseCase{postRepo: postRepo}
}

func (uc *PostUseCase) Create(ctx context.Context, post domain.Post) (*domain.Post, error) {
	if post.UserID <= 0 {
		return nil, domain.ErrIncorrectUserId
	}
	return uc.postRepo.Save(ctx, post)
}

func (uc *PostUseCase) GetAll(ctx context.Context) ([]*domain.Post, error) {
	return uc.postRepo.GetAll(ctx)
}
