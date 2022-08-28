package repository

import (
	"app/entity"
)

type UsersRepository interface {
	Create(params *entity.User) (user *entity.User, err error)
}
