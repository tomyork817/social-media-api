package usecase

import (
	"social-media-api/internal/models"
	"social-media-api/pkg/keys"
	"sync"
)

type SubscriptionUseCase struct {
	lock     sync.RWMutex
	channels map[string]chan *models.Comment
}

func NewSubscriptionUseCase() *SubscriptionUseCase {
	return &SubscriptionUseCase{
		channels: make(map[string]chan *models.Comment),
	}
}

func (uc *SubscriptionUseCase) CreateSubscription(userID, postID int) (<-chan *models.Comment, error) {
	key := keys.GenerateKey(userID, postID)
	uc.lock.Lock()
	defer uc.lock.Unlock()
	if _, exists := uc.channels[key]; !exists {
		uc.channels[key] = make(chan *models.Comment)
		return uc.channels[key], nil
	}
	return nil, models.ErrAlreadyExists
}

func (uc *SubscriptionUseCase) DeleteSubscription(userID, postID int) {
	key := keys.GenerateKey(userID, postID)
	uc.lock.Lock()
	defer uc.lock.Unlock()
	if channel, exists := uc.channels[key]; exists {
		close(channel)
		delete(uc.channels, key)
	}
}

func (uc *SubscriptionUseCase) AddComment(comment *models.Comment) {
	uc.lock.RLock()
	defer uc.lock.RUnlock()
	for key, channel := range uc.channels {
		_, postId := keys.ExtractIds(key)
		if postId == comment.PostID {
			channel <- comment
		}
	}
}
