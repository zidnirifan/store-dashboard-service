package service

import (
	"errors"
	"store-dashboard-service/model"
	"store-dashboard-service/repository/mocks"
	"store-dashboard-service/util/exception"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

var userRepository = new(mocks.UserRepository)
var userServiceTest = NewUserService(userRepository, validator.New())

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
