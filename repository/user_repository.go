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
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (u *userRepository) GetByEmail(email string) (model.User, error) {
	user := model.User{}

	err := u.db.Take(&user, "email = ?", email).Error

	return user, err
}
