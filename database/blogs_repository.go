package database

import (
	"app/entity"
	"app/graphql/model"
)

type BlogsRepository struct {
	DB DB
}

func (repo *BlogsRepository) FindAll(params *model.PageCondition) (blogs []*entity.Blog, err error) {
	db := repo.DB.Connect()

	err = db.Model(&entity.Blog{}).
		Offset((params.PageNo - 1) * params.Limit).
		Limit(params.Limit).
		Find(&blogs).
		Error

	if err != nil {
		return nil, err
	}

	return
}

func (repo *BlogsRepository) Find(id int) (blog *entity.Blog, err error) {
	db := repo.DB.Connect()

	err = db.Model(&entity.Blog{}).Find(&blog, id).Error
	if err != nil {
		return nil, err
	}

	return
}

func (repo *BlogsRepository) TotalCount() (count int64, err error) {
	db := repo.DB.Connect()

	err = db.Model(&entity.Blog{}).Count(&count).Error

	if err != nil {
		return 0, err
	}

	return
}
