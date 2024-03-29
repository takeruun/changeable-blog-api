package resolver

import (
	"app/database"
	"app/service"
	"app/usecases"

	"github.com/wader/gormstore/v2"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserUsecase usecases.UserUsecase
	BlogUsecase usecases.BlogUsecase
}

func NewResolver(db database.DB, store *gormstore.Store) *Resolver {
	sss := service.NewSessionStoreService(store)
	cs := &service.CyptoService{}

	return &Resolver{
		UserUsecase: usecases.UserUsecase{
			UsersRepo:     &database.UsersRepository{DB: db},
			SSService:     sss,
			CryptoService: cs,
		},
		BlogUsecase: usecases.BlogUsecase{
			BlogsRepo: &database.BlogsRepository{DB: db},
		},
	}
}
