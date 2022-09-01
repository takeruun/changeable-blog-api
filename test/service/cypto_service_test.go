package service

import (
	"app/service"
	"testing"
)

func TestHashAndSalt(t *testing.T) {
	password := "password"

	cyptroService := service.CyptoService{}

	hashPassword, err := cyptroService.HashAndSalt([]byte(password))

	if err != nil {
		t.Fatal(err)
		panic("failed hash")
	}

	if password == hashPassword {
		panic("failed hash")
	}
}

func TestComparePasswords(t *testing.T) {
	password := "password"

	cyptroService := service.CyptoService{}

	hashPassword, _ := cyptroService.HashAndSalt([]byte(password))

	compare := cyptroService.ComparePasswords(hashPassword, []byte(password))

	if !compare {
		panic("failed compare")
	}

}
