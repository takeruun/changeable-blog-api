package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/graphql/model"
	"context"
)

// SignUp is the resolver for the signUp field.
func (r *mutationResolver) SignUp(ctx context.Context, input model.SignUp) (*model.User, error) {
	user, err := r.UserUsecase.SignUp(&input, ctx)

	return user, err
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (*model.User, error) {
	user, err := r.UserUsecase.Login(&input, ctx)

	return user, err
}

// GetMyUser is the resolver for the getMyUser field.
func (r *queryResolver) GetMyUser(ctx context.Context) (*model.User, error) {
	user, err := r.UserUsecase.GetMyUser(ctx)

	return user, err
}
