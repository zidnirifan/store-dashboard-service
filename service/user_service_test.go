package service

import (
	"errors"
	"store-dashboard-service/model"
	"store-dashboard-service/repository/mocks"
	"store-dashboard-service/util/exception"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

var userRepository = new(mocks.UserRepository)
var userServiceTest = NewUserService(userRepository, validator.New())

func newUserServiceWithMock() (*userService, *mocks.UserRepository) {
	userRepository := new(mocks.UserRepository)
	userService := NewUserService(userRepository, validator.New())
	return userService, userRepository
}

var passwordHashed, _ = bcrypt.GenerateFromPassword([]byte("password"), 10)
var userFromDB = model.User{
	ID:       "123",
	Email:    "test@mail.com",
	Password: string(passwordHashed),
}

func TestLogin(t *testing.T) {
	t.Run("it should return jwt token", func(t *testing.T) {
		user := model.LoginRequest{
			Email:    "test@mail.com",
			Password: "password",
		}

		userRepository.On("GetByEmail", user.Email).Return(userFromDB, nil)

		res, err := userServiceTest.Login(&user)

		assert.Nil(t, err)
		assert.NotEmpty(t, res.AccessToken)
		assert.NotEmpty(t, res.RefreshToken)
		userRepository.AssertExpectations(t)
	})

	t.Run("it should return validation error when payload invalid", func(t *testing.T) {
		user := model.LoginRequest{
			Email: "user",
		}

		res, err := userServiceTest.Login(&user)

		assert.NotNil(t, err)
		assert.Empty(t, res.AccessToken)
		assert.Empty(t, res.RefreshToken)
		assert.IsType(t, validator.ValidationErrors{}, err)
	})

	t.Run("it should return error 404 when user not found", func(t *testing.T) {
		user := model.LoginRequest{
			Email:    "notfound_user@mail.com",
			Password: "password",
		}

		userRepository.On("GetByEmail", user.Email).Return(model.User{}, errors.New("error"))

		res, err := userServiceTest.Login(&user)

		assert.NotNil(t, err)
		assert.Equal(t, 404, err.(*exception.CustomError).StatusCode)
		assert.Empty(t, res.AccessToken)
		assert.Empty(t, res.RefreshToken)
		userRepository.AssertExpectations(t)
	})

	t.Run("it should return error 403 when password is wrong", func(t *testing.T) {
		user := model.LoginRequest{
			Email:    "user@mail.com",
			Password: "wrong_password",
		}

		userRepository.On("GetByEmail", user.Email).Return(userFromDB, nil)

		res, err := userServiceTest.Login(&user)

		assert.NotNil(t, err)
		assert.Equal(t, 403, err.(*exception.CustomError).StatusCode)
		assert.Empty(t, res.AccessToken)
		assert.Empty(t, res.RefreshToken)
		userRepository.AssertExpectations(t)
	})
}

func TestRegister(t *testing.T) {
	t.Run("it should success register", func(t *testing.T) {
		user := model.RegisterRequest{
			Email:    "test@mail.com",
			Name:     "Test",
			Password: "Password123",
		}

		userService, userRepository := newUserServiceWithMock()
		userRepository.On("GetByEmail", user.Email).Return(model.User{}, errors.New(""))
		userRepository.On("CreateUser", mock.Anything).Return(nil)

		registeredUser, err := userService.Register(&user)

		assert.Nil(t, err)
		assert.NotEmpty(t, registeredUser)
		userRepository.AssertExpectations(t)
	})

	t.Run("it should return validation error when payload is invalid", func(t *testing.T) {
		user := model.RegisterRequest{
			Email: "test",
			Name:  "Test",
		}

		registeredUser, err := userServiceTest.Register(&user)

		assert.NotNil(t, err)
		assert.Empty(t, registeredUser)
		assert.IsType(t, validator.ValidationErrors{}, err)
	})

	t.Run("it should return error 400 when password is not strong", func(t *testing.T) {
		user := model.RegisterRequest{
			Email:    "test@mail.com",
			Name:     "Test",
			Password: "password",
		}

		registeredUser, err := userServiceTest.Register(&user)

		assert.NotNil(t, err)
		assert.Empty(t, registeredUser)
		assert.Equal(t, 400, err.(*exception.CustomError).StatusCode)
	})

	t.Run("it should return error 400 when user already registered", func(t *testing.T) {
		user := model.RegisterRequest{
			Email:    "test@mail.com",
			Name:     "Test",
			Password: "Password123",
		}

		userService, userRepository := newUserServiceWithMock()
		userRepository.On("GetByEmail", user.Email).Return(userFromDB, nil)

		registeredUser, err := userService.Register(&user)

		assert.NotNil(t, err)
		assert.Empty(t, registeredUser)
		assert.Equal(t, 400, err.(*exception.CustomError).StatusCode)
		userRepository.AssertExpectations(t)
	})

	t.Run("it should return error when create user failed", func(t *testing.T) {
		user := model.RegisterRequest{
			Email:    "test@mail.com",
			Name:     "Test",
			Password: "Password123",
		}

		userService, userRepository := newUserServiceWithMock()
		userRepository.On("GetByEmail", user.Email).Return(model.User{}, errors.New(""))
		userRepository.On("CreateUser", mock.Anything).Return(errors.New(""))

		registeredUser, err := userService.Register(&user)

		assert.NotNil(t, err)
		assert.Empty(t, registeredUser)
		userRepository.AssertExpectations(t)
	})
}
