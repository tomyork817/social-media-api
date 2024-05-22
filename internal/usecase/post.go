package usecase

import (
	"context"
	"social-media-api/internal/models"
)

type PostUseCase struct {
	postRepo PostRepo
}

func NewPostUseCase(postRepo PostRepo) *PostUseCase {
	return &PostUseCase{postRepo: postRepo}
}

func (uc *PostUseCase) Create(ctx context.Context, post models.Post) (*models.Post, error) {
	return uc.postRepo.Save(ctx, post)
}

func (uc *PostUseCase) GetAll(ctx context.Context) ([]*models.Post, error) {
	return uc.postRepo.GetAll(ctx)
}
