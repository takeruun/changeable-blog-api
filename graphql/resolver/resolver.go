package resolver

import (
	"app/database"
	"app/interactor"
	"app/service"

	"github.com/wader/gormstore/v2"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UsersInteractor interactor.UsersInteractor
}

func NewResolver(db database.DB, store *gormstore.Store) *Resolver {
	sss := service.NewSessionStoreService(store)
	cs := &service.CyptoService{}

	return &Resolver{
		UsersInteractor: interactor.UsersInteractor{
			UsersRepo:     &database.UsersRepository{DB: db},
			SSService:     sss,
			CryptoService: cs,
		},
	}
}
