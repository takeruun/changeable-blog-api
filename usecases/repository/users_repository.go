package repository

import (
	"app/entity"
)

type UsersRepository interface {
	Create(params *entity.User) (user *entity.User, err error)
	Find(id uint64) (user *entity.User, err error)
	FindByEmail(email string) (user *entity.User, err error)
}
