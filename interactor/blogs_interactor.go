package interactor

import (
	"app/graphql/model"
	"app/interactor/repository"
	"strconv"
)

type BlogsInteractor struct {
	BlogsRepo repository.BlogsRepository
}

func (interactor *BlogsInteractor) BlogList(params *model.PageCondition) (blogListConnection *model.BlogListConnection, err error) {
	blogList, err := interactor.BlogsRepo.FindAll(params)
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

	totalCount, err := interactor.BlogsRepo.TotalCount()

	return &model.BlogListConnection{
		Nodes: node,
		PageInfo: &model.PageInfo{
			TotalCount: int(totalCount),
		},
	}, nil
}

func (interactor *BlogsInteractor) RecommendBlogList() (recommendBlogList *model.RecommendBlogListConnection, err error) {
	blogList, err := interactor.BlogsRepo.FindAll(&model.PageCondition{PageNo: 1, Limit: 3})
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
