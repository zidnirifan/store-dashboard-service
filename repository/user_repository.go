package repository

import (
	"store-dashboard-service/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	GetByEmail(email string) (model.User, error)
	CreateUser(user *model.User) error
	GetById(id string) (model.User, error)
	Update(user model.User) error
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (u *userRepository) GetByEmail(email string) (model.User, error) {
	user := model.User{}

	err := u.db.Take(&user, "email = ?", email).Error

	return user, err
}

func (u *userRepository) CreateUser(user *model.User) error {
	return u.db.Create(&user).Error
}

func (u *userRepository) GetById(id string) (model.User, error) {
	user := model.User{}

	err := u.db.Take(&user, "id = ?", id).Error

	return user, err
}

func (u *userRepository) Update(user model.User) error {
	return u.db.Save(&user).Error
}
