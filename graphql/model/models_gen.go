// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUp struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	PostalCode string `json:"postalCode"`
}

type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	PostalCode string `json:"postalCode"`
	CreatedAt  string `json:"createdAt"`
}
