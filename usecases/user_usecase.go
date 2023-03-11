package usecases

import (
	"app/entity"
	"app/graphql/model"
	"app/service"
	"app/usecases/repository"
	"context"
)

type UserUsecase struct {
	UsersRepo     repository.UsersRepository
	SSService     service.SessionStoreServiceRepository
	CryptoService service.CyptoServiceRepository
}

func (usecase *UserUsecase) SignUp(params *model.SignUp, ctx context.Context) (user *model.User, err error) {
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

func (usecase *UserUsecase) Login(params *model.Login, ctx context.Context) (user *model.User, err error) {
	loginUser, err := usecase.UsersRepo.FindByEmail(params.Email)
	if err != nil {
		return nil, err
	}

	if !usecase.CryptoService.ComparePasswords(loginUser.Password, []byte(params.Password)) {
		return nil, nil
	}

	session, _ := usecase.SSService.GetSession(ctx, "session")

	session.Values["userId"] = loginUser.ID

	usecase.SSService.SaveSession(ctx, session)

	return entity.ToModelUser(loginUser), nil
}

func (usecase *UserUsecase) GetMyUser(ctx context.Context) (user *model.User, err error) {
	session, _ := usecase.SSService.GetSession(ctx, "session")

	userId := session.Values["userId"].(uint64)

	loginUser, err := usecase.UsersRepo.Find(userId)
	if err != nil {
		return nil, err
	}

	return entity.ToModelUser(loginUser), nil
}
