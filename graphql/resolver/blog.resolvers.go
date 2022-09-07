package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/graphql/model"
	"context"
)

// BlogList is the resolver for the blogList field.
func (r *queryResolver) BlogList(ctx context.Context, input model.PageCondition) (*model.BlogListConnection, error) {
	blogListConnection, err := r.BlogsInteractor.BlogList(&input)
	if err != nil {
		return nil, err
	}

	return blogListConnection, nil
}
