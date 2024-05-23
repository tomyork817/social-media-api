package usecase

import (
	"context"
	"social-media-api/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPostRepo struct {
	mock.Mock
}

func (m *MockPostRepo) Save(ctx context.Context, post models.Post) (*models.Post, error) {
	args := m.Called(ctx, post)
	return args.Get(0).(*models.Post), args.Error(1)
}

func (m *MockPostRepo) GetAll(ctx context.Context, limit, offset int) ([]*models.Post, error) {
	args := m.Called(ctx, limit, offset)
	return args.Get(0).([]*models.Post), args.Error(1)
}

func (m *MockPostRepo) GetByUserID(ctx context.Context, userID int, limit, offset int) ([]*models.Post, error) {
	args := m.Called(ctx, userID, limit, offset)
	return args.Get(0).([]*models.Post), args.Error(1)
}

func (m *MockPostRepo) GetByID(ctx context.Context, id int) (*models.Post, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Post), args.Error(1)
}

func (m *MockPostRepo) UpdateIsOpenById(ctx context.Context, id int, isOpen bool) (*models.Post, error) {
	args := m.Called(ctx, id, isOpen)
	return args.Get(0).(*models.Post), args.Error(1)
}

func TestPostUseCaseCreate(t *testing.T) {
	mockPostRepo := new(MockPostRepo)
	mockPostRepo.On("Save", mock.Anything, mock.AnythingOfType("models.Post")).
		Return(&models.Post{ID: 1, UserID: 1, Body: "Test post", IsOpen: true}, nil)

	postUseCase := NewPostUseCase(mockPostRepo)
	post, err := postUseCase.Create(context.Background(), models.Post{UserID: 1, Body: "Test post", IsOpen: true})

	assert.NoError(t, err)
	assert.Equal(t, 1, post.ID)
	assert.Equal(t, 1, post.UserID)
	assert.Equal(t, "Test post", post.Body)
	assert.True(t, post.IsOpen)
	mockPostRepo.AssertExpectations(t)
}

func TestPostUseCaseGetMultiple(t *testing.T) {
	mockPostRepo := new(MockPostRepo)
	mockPostRepo.On("GetAll", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).
		Return([]*models.Post{{ID: 1, UserID: 1, Body: "Test post 1", IsOpen: true}, {ID: 2, UserID: 1, Body: "Test post 2", IsOpen: true}}, nil)

	postUseCase := NewPostUseCase(mockPostRepo)
	posts, err := postUseCase.GetMultiple(context.Background(), models.PostFilter{}, 2, 0)

	assert.NoError(t, err)
	assert.Len(t, posts, 2)
	assert.Equal(t, 1, posts[0].ID)
	assert.Equal(t, 1, posts[0].UserID)
	assert.Equal(t, "Test post 1", posts[0].Body)
	assert.True(t, posts[0].IsOpen)
	assert.Equal(t, 2, posts[1].ID)
	assert.Equal(t, 1, posts[1].UserID)
	assert.Equal(t, "Test post 2", posts[1].Body)
	assert.True(t, posts[1].IsOpen)
	mockPostRepo.AssertExpectations(t)
}

func TestPostUseCaseGetById(t *testing.T) {
	mockPostRepo := new(MockPostRepo)
	mockPostRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int")).
		Return(&models.Post{ID: 1, UserID: 1, Body: "Test post", IsOpen: true}, nil)

	postUseCase := NewPostUseCase(mockPostRepo)
	post, err := postUseCase.GetById(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, 1, post.ID)
	assert.Equal(t, 1, post.UserID)
	assert.Equal(t, "Test post", post.Body)
	assert.True(t, post.IsOpen)
	mockPostRepo.AssertExpectations(t)
}

func TestPostUseCaseDisableComments(t *testing.T) {
	mockPostRepo := new(MockPostRepo)
	mockPostRepo.On("UpdateIsOpenById", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("bool")).
		Return(&models.Post{ID: 1, UserID: 1, Body: "Test post", IsOpen: false}, nil)

	postUseCase := NewPostUseCase(mockPostRepo)
	post, err := postUseCase.DisableComments(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, 1, post.ID)
	assert.Equal(t, 1, post.UserID)
	assert.Equal(t, "Test post", post.Body)
	assert.False(t, post.IsOpen)
	mockPostRepo.AssertExpectations(t)
}

func TestPostUseCaseEnableComments(t *testing.T) {
	mockPostRepo := new(MockPostRepo)
	mockPostRepo.On("UpdateIsOpenById", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("bool")).
		Return(&models.Post{ID: 1, UserID: 1, Body: "Test post", IsOpen: true}, nil)

	postUseCase := NewPostUseCase(mockPostRepo)
	post, err := postUseCase.EnableComments(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, 1, post.ID)
	assert.Equal(t, 1, post.UserID)
	assert.Equal(t, "Test post", post.Body)
	assert.True(t, post.IsOpen)
	mockPostRepo.AssertExpectations(t)
}
