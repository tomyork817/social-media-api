package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"social-media-api/internal/controller/graph/generated"
	"social-media-api/internal/models"
)

// NewComment is the resolver for the newComment field.
func (r *subscriptionResolver) NewComment(ctx context.Context, userID int, postID int) (<-chan *models.Comment, error) {
	ch, err := r.SubscriptionUseCase.CreateSubscription(userID, postID)
	if err != nil {
		return nil, err
	}
	go func() {
		defer r.SubscriptionUseCase.DeleteSubscription(userID, postID)
		for {
			select {
			case <-ctx.Done():
				return
			}
		}
	}()

	return ch, nil

}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }