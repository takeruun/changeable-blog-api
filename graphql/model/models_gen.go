// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type LoginUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SignUp struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	PostalCode string `json:"postalCode"`
}