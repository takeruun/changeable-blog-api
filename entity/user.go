package entity

import (
	"app/graphql/model"
	"strconv"
	"time"
)

type User struct {
	ID         uint64 `json:"id" gorm:"primary_key"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string
	PostalCode string     `json:"postal_code"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt"`
}

func ToEntityUser(u *model.SignUp) *User {
	return &User{
		Name:       u.Name,
		Email:      u.Email,
		Password:   u.Password,
		PostalCode: u.PostalCode,
	}
}

func ToModelUser(u *User) *model.User {
	return &model.User{
		ID:         strconv.Itoa(int(u.ID)),
		Name:       u.Name,
		PostalCode: u.PostalCode,
	}
}
