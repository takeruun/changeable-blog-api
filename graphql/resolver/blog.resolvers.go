package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/graphql/model"
	"context"
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

// BlogList is the resolver for the blogList field.
func (r *queryResolver) BlogList(ctx context.Context, input model.PageCondition) (*model.BlogListConnection, error) {
	blogListConnection, err := r.BlogUsecase.BlogList(&input)
	if err != nil {
		return nil, err
	}

	return blogListConnection, nil
}

// RecommendBlogList is the resolver for the recommendBlogList field.
func (r *queryResolver) RecommendBlogList(ctx context.Context) (*model.RecommendBlogListConnection, error) {
	recommendBlogList, err := r.BlogUsecase.RecommendBlogList()
	if err != nil {
		return nil, err
	}

	return recommendBlogList, nil
}

// Blog is the resolver for the blog field.
func (r *queryResolver) Blog(ctx context.Context, id int) (*model.Blog, error) {
	blog, err := r.BlogUsecase.Blog(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		graphql.AddError(ctx, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: fmt.Sprintf("Error %s", err),
			Extensions: map[string]interface{}{
				"code": "NOT_FOUND_ERROR",
			},
		})
	} else if err != nil {
		return nil, err
	}

	return blog, nil
}
