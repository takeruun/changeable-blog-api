package usecases_test

import (
	"app/entity"
	"app/graphql/model"
	"app/test/mock_repository"
	"app/test/mock_service"
	"app/usecases"
	"context"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
)

var mockUsersRepository *mock_repository.MockUsersRepository
var mockSessionStoreService *mock_service.MockSessionStoreServiceRepository
var mockCyptroService *mock_service.MockCyptoServiceRepository
var userUsecase usecases.UserUsecase

func userTestSetUp(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsersRepository = mock_repository.NewMockUsersRepository(ctrl)
	mockSessionStoreService = mock_service.NewMockSessionStoreServiceRepository(ctrl)
	mockCyptroService = mock_service.NewMockCyptoServiceRepository(ctrl)

	userUsecase = usecases.UserUsecase{
		UsersRepo:     mockUsersRepository,
		SSService:     mockSessionStoreService,
		CryptoService: mockCyptroService,
	}

	return func() {
		defer ctrl.Finish()
	}
}

func TestSignUp(t *testing.T) {
	userTestSetUp := userTestSetUp(t)
	defer userTestSetUp()

	var (
		ID        uint64 = 1
		name             = "test"
		email            = "test@example.com"
		posalCode        = "100-0001"
		password         = "password"
		err       error  = nil
	)
	user := entity.User{Name: name, Email: email, Password: password, PostalCode: posalCode}
	newUser := entity.User{ID: ID, Name: name, Email: email, Password: password, PostalCode: posalCode}
	expected := model.User{ID: strconv.Itoa(int(ID)), Name: name}

	mockCyptroService.EXPECT().HashAndSalt([]byte(password)).Return(password, nil)
	mockUsersRepository.EXPECT().Create(&user).Return(&newUser, err)
	mockSessionStoreService.EXPECT().SaveValue(gomock.Any(), "userId", newUser.ID)

	params := &model.SignUp{Name: name, Email: email, Password: password, PostalCode: posalCode}

	modelUser, err := userUsecase.SignUp(params, context.TODO())

	assert.NoError(t, err)

	assert.Equal(t, expected.ID, modelUser.ID)
	assert.Equal(t, expected.Name, modelUser.Name)
}

func TestLogin(t *testing.T) {
	userTestSetUp := userTestSetUp(t)
	defer userTestSetUp()

	var (
		ID        uint64 = 1
		name             = "test"
		email            = "test@example.com"
		posalCode        = "100-0001"
		password         = "password"
	)
	loginUser := entity.User{ID: ID, Name: name, Email: email, Password: password, PostalCode: posalCode}
	mockSession := &sessions.Session{
		Values: map[interface{}]interface{}{"userId": "1"},
	}
	expected := model.User{ID: strconv.Itoa(int(ID)), Name: name}

	mockUsersRepository.EXPECT().FindByEmail(email).Return(&loginUser, nil)
	mockCyptroService.EXPECT().ComparePasswords(password, []byte(password)).Return(true)
	mockSessionStoreService.EXPECT().GetSession(gomock.Any(), "session").Return(mockSession, nil)
	mockSessionStoreService.EXPECT().SaveSession(gomock.Any(), mockSession)

	params := &model.Login{Password: password, Email: email}
	modelUser, err := userUsecase.Login(params, context.TODO())

	assert.NoError(t, err)

	assert.Equal(t, expected.ID, modelUser.ID)
	assert.Equal(t, expected.Name, modelUser.Name)
}

func TestGetMyUser(t *testing.T) {
	userTestSetUp := userTestSetUp(t)
	defer userTestSetUp()

	var (
		userId    uint64 = 1
		name             = "test"
		email            = "test@example.com"
		posalCode        = "100-0001"
		password         = "password"
	)
	mockSession := &sessions.Session{
		Values: map[interface{}]interface{}{"userId": userId},
	}
	loginUser := entity.User{ID: userId, Name: name, Email: email, Password: password, PostalCode: posalCode}
	expected := model.User{ID: strconv.Itoa(int(userId)), Name: name, PostalCode: posalCode}

	mockSessionStoreService.EXPECT().GetSession(gomock.Any(), "session").Return(mockSession, nil)
	mockUsersRepository.EXPECT().Find(userId).Return(&loginUser, nil)

	modelUser, err := userUsecase.GetMyUser(context.TODO())

	assert.NoError(t, err)

	assert.Equal(t, expected.ID, modelUser.ID)
	assert.Equal(t, expected.Name, modelUser.Name)
}
