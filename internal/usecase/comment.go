package usecase

import (
	"context"
	"social-media-api/internal/models"
)

type CommentUseCase struct {
	commentRepo CommentRepo
	postRepo    PostRepo
}

func NewCommentUseCase(commentRepo CommentRepo, postRepo PostRepo) *CommentUseCase {
	return &CommentUseCase{commentRepo: commentRepo, postRepo: postRepo}
}

func (uc *CommentUseCase) Create(ctx context.Context, comment models.Comment) (*models.Comment, error) {
	postID := comment.PostID
	post, _ := uc.postRepo.GetByID(ctx, postID)
	if post == nil {
		return nil, models.ErrNotFound
	}
	if !post.IsOpen {
		return nil, models.ErrPostCommentsDisabled
	}
	if comment.ParentID != 0 {
		parent, _ := uc.commentRepo.GetByID(ctx, comment.ParentID)
		if parent == nil || parent.PostID != postID {
			return nil, models.ErrNotFound
		}
	}
	return uc.commentRepo.Save(ctx, comment)
}

func (uc *CommentUseCase) GetMultiple(ctx context.Context, filter models.CommentFilter, limit, offset int) ([]*models.Comment, error) {
	if filter.PostID != 0 {
		return uc.commentRepo.GetByPostID(ctx, filter.PostID, limit, offset)
	}
	if filter.ParentID != 0 {
		return uc.commentRepo.GetByParentID(ctx, filter.ParentID, limit, offset)
	}
	return nil, models.ErrIncorrectFilter
}

func (uc *CommentUseCase) GetById(ctx context.Context, id int) (*models.Comment, error) {
	return uc.GetById(ctx, id)
}
