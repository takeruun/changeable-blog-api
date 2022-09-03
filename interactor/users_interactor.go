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

func (usecase *UsersInteractor) SignUp(params *model.SignUp, ctx context.Context) (user *model.User, err error) {
	hashPwd, err := usecase.CryptoService.HashAndSalt([]byte(params.Password))
	if err != nil {
		return nil, err
	}

	params.Password = hashPwd
	userParams := entity.ToEntityUser(params)

	newUser, err := usecase.UsersRepo.Create(userParams)
	if err != nil {
		return nil, err
	}

	usecase.SSService.SaveValue(ctx, "userId", newUser.ID)

	return entity.ToModelUser(newUser), nil
}

func (usecase *UsersInteractor) Login(params *model.Login, ctx context.Context) (user *model.User, err error) {

	loginUser, err := usecase.UsersRepo.FindByEmail(params.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(loginUser.Password), []byte(params.Password))
	if err != nil {
		return nil, err
	}

	session, _ := usecase.SSService.GetSession(ctx, "session")

	session.Values["userId"] = loginUser.ID

	usecase.SSService.SaveSession(ctx, session)

	return entity.ToModelUser(loginUser), nil
}

func (usecase *UsersInteractor) GetMyUser(ctx context.Context) (user *model.User, err error) {
	session, _ := usecase.SSService.GetSession(ctx, "session")

	userId := session.Values["userId"].(uint64)

	loginUser, err := usecase.UsersRepo.Find(userId)
	if err != nil {
		return nil, err
	}

	return entity.ToModelUser(loginUser), nil
}
