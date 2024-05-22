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

func (uc *PostUseCase) GetMultiple(ctx context.Context, filter models.PostFilter, limit, offset int) ([]*models.Post, error) {
	if filter.UserID == 0 {
		return uc.postRepo.GetAll(ctx, limit, offset)
	}
	return uc.postRepo.GetByUserID(ctx, filter.UserID, limit, offset)
}

func (uc *PostUseCase) GetById(ctx context.Context, id int) (*models.Post, error) {
	return uc.postRepo.GetByID(ctx, id)
}

func (uc *PostUseCase) DisableComments(ctx context.Context, id int) (*models.Post, error) {
	return uc.postRepo.UpdateIsOpenById(ctx, id, false)
}

func (uc *PostUseCase) EnableComments(ctx context.Context, id int) (*models.Post, error) {
	return uc.postRepo.UpdateIsOpenById(ctx, id, true)
}
