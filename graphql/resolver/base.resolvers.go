package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/graphql/generated"
	"context"
	"fmt"
)

// Empty is the resolver for the _empty field.
func (r *mutationResolver) Empty(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Empty is the resolver for the _empty field.
func (r *queryResolver) Empty(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
