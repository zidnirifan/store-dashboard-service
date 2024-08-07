package service

import (
	"errors"
	"fmt"
	"store-dashboard-service/model"
	"store-dashboard-service/repository"
	"store-dashboard-service/util/exception"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repository repository.UserRepository
}

type UserService interface {
	Login(payload *model.LoginRequest) (model.Token, error)
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository: repository}
}

func (u *userService) Login(payload *model.LoginRequest) (model.Token, error) {
	token := model.Token{}
	user, err := u.repository.GetByEmail(payload.Email)
	if err != nil {
		return token, &exception.CustomError{StatusCode: 404, Err: errors.New("user not found")}
	}

	passwordMatchErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if passwordMatchErr != nil {
		return token, &exception.CustomError{StatusCode: 403, Err: errors.New("wrong password")}
	}

	fmt.Println(user)

	// TODO replace with real jwt token
	token.AccessToken = "access token"
	token.RefreshToken = "refresh token"
	return token, nil
}
