package mock_repository

import (
	"app/entity"
	"app/graphql/generated"
	"app/graphql/model"
	"app/graphql/resolver"
	"app/test/mock_service"
	"app/usecases"
	"strconv"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var mockUsersRepository *MockUsersRepository
var mockSessionStoreService *mock_service.MockSessionStoreServiceRepository
var mockCyptroService *mock_service.MockCyptoServiceRepository
var resolvers resolver.Resolver

func setUp(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsersRepository = NewMockUsersRepository(ctrl)
	mockSessionStoreService = mock_service.NewMockSessionStoreServiceRepository(ctrl)
	mockCyptroService = mock_service.NewMockCyptoServiceRepository(ctrl)

	usecase := usecases.UserUsecase{
		UsersRepo:     mockUsersRepository,
		SSService:     mockSessionStoreService,
		CryptoService: mockCyptroService,
	}

	resolvers = resolver.Resolver{UserUsecase: usecase}

	return func() {
		defer ctrl.Finish()
	}
}

func TestSignUp(t *testing.T) {
	setup := setUp(t)
	defer setup()

	var (
		ID        uint64 = 1
		name             = "test"
		email            = "test@example.com"
		posalCode        = "100-0001"
		password         = "password"
		err       error  = nil
	)
	params := entity.User{Name: name, Email: email, Password: password, PostalCode: posalCode}
	newUser := entity.User{ID: ID, Name: name, Email: email, Password: password, PostalCode: posalCode}
	modelUser := model.User{ID: strconv.Itoa(int(ID)), Name: name}

	mockCyptroService.EXPECT().HashAndSalt([]byte(password)).Return(password, nil)
	mockUsersRepository.EXPECT().Create(&params).Return(&newUser, err)
	mockSessionStoreService.EXPECT().SaveValue(gomock.Any(), "userId", newUser.ID)

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers})))

	var resp struct {
		SignUp struct {
			ID, Name string
		}
	}

	query := `
		mutation {
			signUp(
				input: {name: "test", email: "test@example.com", password: "password", postalCode: "100-0001"}
			) {
				name
				id
			}
		}
	`

	c.MustPost(query, &resp)
	assert.Equal(t, modelUser.ID, resp.SignUp.ID)
	assert.Equal(t, modelUser.Name, resp.SignUp.Name)
}
