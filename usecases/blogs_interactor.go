package usecases

import (
	"app/graphql/model"
	"app/usecases/repository"
	"strconv"
)

type BlogsUsecases struct {
	BlogsRepo repository.BlogsRepository
}

func (usecases *BlogsUsecases) BlogList(params *model.PageCondition) (blogListConnection *model.BlogListConnection, err error) {
	blogList, err := usecases.BlogsRepo.FindAll(params)
	if err != nil {
		return nil, err
	}

	var node []*model.BlogList
	for _, blog := range blogList {
		node = append(node, &model.BlogList{
			ID:                 strconv.Itoa(int(blog.ID)),
			Title:              blog.Title,
			ThumbnailImagePath: blog.ThumbnailImagePath,
			CreatedAt:          blog.CreatedAt.Format("2006-01-02"),
		})
	}

	totalCount, err := usecases.BlogsRepo.TotalCount()

	return &model.BlogListConnection{
		Nodes: node,
		PageInfo: &model.PageInfo{
			TotalCount: int(totalCount),
		},
	}, nil
}

func (usecases *BlogsUsecases) RecommendBlogList() (recommendBlogList *model.RecommendBlogListConnection, err error) {
	blogList, err := usecases.BlogsRepo.FindAll(&model.PageCondition{PageNo: 0, Limit: 3})
	if err != nil {
		return nil, err
	}

	var node []*model.BlogList
	for _, blog := range blogList {
		node = append(node, &model.BlogList{
			ID:                 strconv.Itoa(int(blog.ID)),
			Title:              blog.Title,
			ThumbnailImagePath: blog.ThumbnailImagePath,
			CreatedAt:          blog.CreatedAt.Format("2006-01-02"),
		})
	}

	return &model.RecommendBlogListConnection{
		Nodes: node,
	}, nil
}

func (usecases *BlogsUsecases) Blog(id int) (modelaBlog *model.Blog, err error) {
	blog, err := usecases.BlogsRepo.Find(id)
	if err != nil {
		return nil, err
	}

	return &model.Blog{
		ID:                 strconv.Itoa(int(blog.ID)),
		Body:               blog.Body,
		NightBody:          blog.NgihtBody,
		MobileBody:         blog.MobileBody,
		Title:              blog.Title,
		ThumbnailImagePath: blog.ThumbnailImagePath,
		Tags:               []string{"sample"},
		CreatedAt:          blog.CreatedAt.Format("2006-01-02"),
		UpdatedAt:          blog.UpdatedAt.Format("2006-01-02"),
	}, nil
}
