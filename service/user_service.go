package service

import (
	"errors"
	"store-dashboard-service/model"
	"store-dashboard-service/repository"
	"store-dashboard-service/util"
	"store-dashboard-service/util/exception"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repository repository.UserRepository
	validate   *validator.Validate
}

type UserService interface {
	Login(payload *model.LoginRequest) (model.Token, error)
	Register(payload *model.RegisterRequest) (model.RegisteredUser, error)
	VerifyUser(userId string) error
}

func NewUserService(repository repository.UserRepository, validate *validator.Validate) *userService {
	return &userService{repository: repository, validate: validate}
}

func (u *userService) Login(payload *model.LoginRequest) (model.Token, error) {
	token := model.Token{}

	err := u.validate.Struct(payload)
	if err != nil {
		return token, err
	}

	user, err := u.repository.GetByEmail(payload.Email)
	if err != nil {
		return token, &exception.CustomError{StatusCode: 404, Err: errors.New("user not found")}
	}
	if user.Status != util.CommonConst.Status.Active {
		return token, &exception.CustomError{StatusCode: 403, Err: errors.New("user not activated")}
	}

	passwordMatchErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if passwordMatchErr != nil {
		return token, &exception.CustomError{StatusCode: 403, Err: errors.New("wrong password")}
	}

	accessToken, err := util.GenerateAccessToken(&model.PayloadAccessToken{ID: user.ID, Email: user.Email, Role: user.Role})
	refreshToken, err := util.GenerateRefreshToken(&model.PayloadRefreshToken{ID: user.ID, Email: user.Email})
	if err != nil {
		return token, err
	}

	token.AccessToken = accessToken
	token.RefreshToken = refreshToken

	return token, nil
}

func (u *userService) Register(payload *model.RegisterRequest) (model.RegisteredUser, error) {
	registeredUser := model.RegisteredUser{}

	err := u.validate.Struct(payload)
	if err != nil {
		return registeredUser, err
	}

	matched := util.ValidatePassword(payload.Password)
	if !matched {
		errMsg := "password must be at least 8 characters, one number, one uppercase letter, and one lowercase letter"
		return registeredUser, &exception.CustomError{StatusCode: 400, Err: errors.New(errMsg)}
	}

	passwordHashed, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)
	user := model.User{
		ID:       uuid.NewString(),
		Name:     payload.Name,
		Email:    payload.Email,
		Password: string(passwordHashed),
		Role:     util.CommonConst.Roles.Admin,
		Status:   util.CommonConst.Status.NotVerified,
	}

	_, err = u.repository.GetByEmail(payload.Email)
	if err == nil {
		return registeredUser, &exception.CustomError{StatusCode: 400, Err: errors.New("user already registered")}
	}

	err = u.repository.CreateUser(&user)
	if err != nil {
		return registeredUser, err
	}

	registeredUser = model.RegisteredUser{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
	}

	return registeredUser, nil
}

func (u *userService) VerifyUser(userId string) error {
	user, err := u.repository.GetById(userId)
	if err != nil {
		return &exception.CustomError{StatusCode: 404, Err: errors.New("user not found")}
	}

	if user.Status == util.CommonConst.Status.Active {
		return &exception.CustomError{StatusCode: 400, Err: errors.New("user already verified")}
	}

	user.Status = util.CommonConst.Status.Active
	err = u.repository.Update(&user)

	return err
}
