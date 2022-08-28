package usecase

import (
	"app/entity"
	"app/graphql/model"
	"app/service"
	"app/usecase/repository"
	"context"

	"golang.org/x/crypto/bcrypt"
)

type UsersUsecase struct {
	UsersRepo repository.UsersRepository
	SSService service.SessionStoreService
}

func (usecase *UsersUsecase) SignUp(params *model.SignUp, ctx context.Context) (user *model.User, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	params.Password = string(hash)
	userParams := entity.ToEntityUser(params)

	newUser, err := usecase.UsersRepo.Create(userParams)
	if err != nil {
		return nil, err
	}

	session, _ := usecase.SSService.GetSession(ctx, "session")

	session.Values["userId"] = newUser.ID

	usecase.SSService.SaveSession(ctx, session)

	return entity.ToModelUser(newUser), nil
}

func (usecase *UsersUsecase) Login(params *model.Login, ctx context.Context) (user *model.User, err error) {

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

func (usecase *UsersUsecase) GetMyUser(ctx context.Context) (user *model.User, err error) {
	session, _ := usecase.SSService.GetSession(ctx, "session")

	userId := session.Values["userId"].(uint64)

	loginUser, err := usecase.UsersRepo.Find(userId)
	if err != nil {
		return nil, err
	}

	return entity.ToModelUser(loginUser), nil
}
