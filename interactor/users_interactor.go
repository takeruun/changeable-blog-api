package interactor

import (
	"app/entity"
	"app/graphql/model"
	"app/interactor/repository"
	"app/service"
	"context"

	"golang.org/x/crypto/bcrypt"
)

type UsersInteractor struct {
	UsersRepo     repository.UsersRepository
	SSService     service.SessionStoreServiceRepository
	CryptoService service.CyptoServiceRepository
}

func (interactor *UsersInteractor) SignUp(params *model.SignUp, ctx context.Context) (user *model.User, err error) {
	hashPwd, err := interactor.CryptoService.HashAndSalt([]byte(params.Password))
	if err != nil {
		return nil, err
	}

	params.Password = hashPwd
	userParams := entity.ToEntityUser(params)

	newUser, err := interactor.UsersRepo.Create(userParams)
	if err != nil {
		return nil, err
	}

	interactor.SSService.SaveValue(ctx, "userId", newUser.ID)

	return entity.ToModelUser(newUser), nil
}

func (interactor *UsersInteractor) Login(params *model.Login, ctx context.Context) (user *model.User, err error) {
	loginUser, err := interactor.UsersRepo.FindByEmail(params.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(loginUser.Password), []byte(params.Password))
	if err != nil {
		return nil, err
	}

	session, _ := interactor.SSService.GetSession(ctx, "session")

	session.Values["userId"] = loginUser.ID

	interactor.SSService.SaveSession(ctx, session)

	return entity.ToModelUser(loginUser), nil
}

func (interactor *UsersInteractor) GetMyUser(ctx context.Context) (user *model.User, err error) {
	session, _ := interactor.SSService.GetSession(ctx, "session")

	userId := session.Values["userId"].(uint64)

	loginUser, err := interactor.UsersRepo.Find(userId)
	if err != nil {
		return nil, err
	}

	return entity.ToModelUser(loginUser), nil
}
