package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.21-dev DO NOT EDIT.

import (
	"context"

	"github.com/99designs/gqlgen/_examples/federation/reviews/graph/model"
)

// Reviews is the resolver for the reviews field.
func (r *userResolver) Reviews(ctx context.Context, obj *model.User) ([]*model.Review, error) {
	var productReviews []*model.Review
	for _, review := range reviews {
		if review.Author.ID == obj.ID {
			productReviews = append(productReviews, review)
		}
	}
	return productReviews, nil
}

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
