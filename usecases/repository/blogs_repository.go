package repository

import (
	"app/entity"
	"app/graphql/model"
)

type BlogsRepository interface {
	FindAll(params *model.PageCondition) (blogs []*entity.Blog, err error)
	TotalCount() (count int64, err error)
	Find(id int) (blog *entity.Blog, err error)
}
