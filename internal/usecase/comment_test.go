package usecase

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"social-media-api/internal/models"
	"testing"
)

type MockCommentRepo struct {
	mock.Mock
}

func (m *MockCommentRepo) Save(ctx context.Context, comment models.Comment) (*models.Comment, error) {
	args := m.Called(ctx, comment)
	return args.Get(0).(*models.Comment), args.Error(1)
}

func (m *MockCommentRepo) GetByID(ctx context.Context, id int) (*models.Comment, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Comment), args.Error(1)
}

func (m *MockCommentRepo) GetByPostID(ctx context.Context, postID int, limit, offset int) ([]*models.Comment, error) {
	args := m.Called(ctx, postID, limit, offset)
	return args.Get(0).([]*models.Comment), args.Error(1)
}

func (m *MockCommentRepo) GetByParentID(ctx context.Context, parentID int, limit, offset int) ([]*models.Comment, error) {
	args := m.Called(ctx, parentID, limit, offset)
	return args.Get(0).([]*models.Comment), args.Error(1)
}

func (m *MockCommentRepo) GetAll(ctx context.Context, limit, offset int) ([]*models.Comment, error) {
	args := m.Called(ctx, limit, offset)
	return args.Get(0).([]*models.Comment), args.Error(1)
}

func TestCommentUseCaseCreate(t *testing.T) {
	mockCommentRepo := new(MockCommentRepo)
	mockPostRepo := new(MockPostRepo)

	mockPostRepo.On("GetByID", mock.Anything, 1).
		Return(&models.Post{ID: 1, UserID: 1, Body: "Test post", IsOpen: true}, nil)

	mockCommentRepo.On("Save", mock.Anything, mock.AnythingOfType("models.Comment")).
		Return(&models.Comment{ID: 1, UserID: 1, PostID: 1, Body: "Test comment"}, nil)

	commentUseCase := NewCommentUseCase(mockCommentRepo, mockPostRepo)
	comment, err := commentUseCase.Create(context.Background(), models.Comment{UserID: 1, PostID: 1, Body: "Test comment"})

	assert.NoError(t, err)
	assert.Equal(t, 1, comment.ID)
	assert.Equal(t, 1, comment.UserID)
	assert.Equal(t, 1, comment.PostID)
	assert.Equal(t, "Test comment", comment.Body)
	mockPostRepo.AssertExpectations(t)
	mockCommentRepo.AssertExpectations(t)
}

func TestCommentUseCaseGetMultiple(t *testing.T) {
	mockCommentRepo := new(MockCommentRepo)
	mockPostRepo := new(MockPostRepo)

	mockCommentRepo.On("GetByPostID", mock.Anything, 1, 2, 0).
		Return([]*models.Comment{
			{ID: 1, UserID: 1, PostID: 1, Body: "Test comment 1"},
			{ID: 2, UserID: 1, PostID: 1, Body: "Test comment 2"},
		}, nil)

	commentUseCase := NewCommentUseCase(mockCommentRepo, mockPostRepo)
	comments, err := commentUseCase.GetMultiple(context.Background(), models.CommentFilter{PostID: 1}, 2, 0)

	assert.NoError(t, err)
	assert.Len(t, comments, 2)
	assert.Equal(t, 1, comments[0].ID)
	assert.Equal(t, 1, comments[0].UserID)
	assert.Equal(t, 1, comments[0].PostID)
	assert.Equal(t, "Test comment 1", comments[0].Body)
	assert.Equal(t, 2, comments[1].ID)
	assert.Equal(t, 1, comments[1].UserID)
	assert.Equal(t, 1, comments[1].PostID)
	assert.Equal(t, "Test comment 2", comments[1].Body)
	mockCommentRepo.AssertExpectations(t)
}
func TestCommentUseCaseGetById(t *testing.T) {
	mockCommentRepo := new(MockCommentRepo)
	mockPostRepo := new(MockPostRepo)

	mockCommentRepo.On("GetByID", mock.Anything, 1).
		Return(&models.Comment{ID: 1, UserID: 1, PostID: 1, Body: "Test comment"}, nil)

	commentUseCase := NewCommentUseCase(mockCommentRepo, mockPostRepo)
	comment, err := commentUseCase.GetById(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, 1, comment.ID)
	assert.Equal(t, 1, comment.UserID)
	assert.Equal(t, 1, comment.PostID)
	assert.Equal(t, "Test comment", comment.Body)
	mockCommentRepo.AssertExpectations(t)
}
