package service

import (
	"fmt"
	"store-dashboard-service/model"
	"store-dashboard-service/repository"

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
		return token, err
	}

	passwordMatchErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if passwordMatchErr != nil {
		return token, err
	}

	fmt.Println(user)

	// TODO replace with real jwt token
	token.AccessToken = "access token"
	token.RefreshToken = "refresh token"
	return token, nil
}
