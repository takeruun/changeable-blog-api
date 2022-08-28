package usecase

import (
	"app/entity"
	"app/graphql/model"
	"app/service"
	"app/usecase/repository"
	"context"
)

type UsersUsecase struct {
	UsersRepo repository.UsersRepository
	SSService service.SessionStoreService
}

func (usecase *UsersUsecase) SignUp(params *model.SignUp, ctx context.Context) (loginUser *model.LoginUser, err error) {
	userParams := entity.ToEntityUser(params)

	user, err := usecase.UsersRepo.Create(userParams)
	if err != nil {
		return nil, err
	}

	session, _ := usecase.SSService.GetSession(ctx, "session")

	session.Values["userId"] = user.ID

	usecase.SSService.SaveSession(ctx, session)

	return entity.ToModelLoginUser(user), nil
}
