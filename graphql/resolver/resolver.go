package resolver

import (
	"app/database"
	"app/service"
	"app/usecase"

	"github.com/wader/gormstore/v2"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UsecaseUsers usecase.UsersUsecase
}

func NewResolver(db database.DB, store *gormstore.Store) *Resolver {
	sss := service.NewSessionStoreService(store)
	cs := &service.CyptoService{}

	return &Resolver{
		UsecaseUsers: usecase.UsersUsecase{
			UsersRepo:     &database.UsersRepository{DB: db},
			SSService:     sss,
			CryptoService: cs,
		},
	}
}
