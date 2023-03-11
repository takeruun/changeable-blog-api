package database

import (
	"app/entity"
)

type UsersRepository struct {
	DB DB
}

func (repo *UsersRepository) Create(params *entity.User) (user *entity.User, err error) {
	db := repo.DB.Connect()

	result := db.Create(&params)
	if result.Error != nil {
		return &entity.User{}, result.Error
	}

	result = db.Find(&user, params.ID)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repo *UsersRepository) Find(id uint64) (user *entity.User, err error) {
	db := repo.DB.Connect()

	result := db.Find(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repo *UsersRepository) FindByEmail(email string) (user *entity.User, err error) {
	db := repo.DB.Connect()

	result := db.Find(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
