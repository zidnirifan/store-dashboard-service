package repository

import (
	"store-dashboard-service/model"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *model.Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

func (c *categoryRepository) Create(category *model.Category) error {
	return c.db.Create(category).Error
}
