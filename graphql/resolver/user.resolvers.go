package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/graphql/model"
	"context"
)

// SignUp is the resolver for the signUp field.
func (r *mutationResolver) SignUp(ctx context.Context, input model.SignUp) (*model.User, error) {
	user, _ := r.uc.SignUp(&input, ctx)

	return user, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (*model.User, error) {
	user, _ := r.uc.Login(&input, ctx)

	return user, nil
}

// GetMyUser is the resolver for the getMyUser field.
func (r *queryResolver) GetMyUser(ctx context.Context) (*model.User, error) {
	user, _ := r.uc.GetMyUser(ctx)

	return user, nil
}
