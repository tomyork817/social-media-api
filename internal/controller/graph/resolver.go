package graph

import (
	"social-media-api/internal/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PostUseCase         usecase.Post
	CommentUseCase      usecase.Comment
	SubscriptionUseCase usecase.Subscription
}

func NewResolver(postUseCase usecase.Post, commentUseCase usecase.Comment, subscriptionUseCase usecase.Subscription) *Resolver {
	return &Resolver{PostUseCase: postUseCase, CommentUseCase: commentUseCase, SubscriptionUseCase: subscriptionUseCase}
}
