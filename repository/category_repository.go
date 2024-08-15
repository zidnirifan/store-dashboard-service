package repository

import (
	"store-dashboard-service/model"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *model.Category) error
	GetAll(categories *[]model.Category) error
	GetById(id int) (model.Category, error)
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

func (c *categoryRepository) GetAll(categories *[]model.Category) error {
	return c.db.Find(categories).Error
}

func (c *categoryRepository) GetById(id int) (model.Category, error) {
	category := model.Category{}

	err := c.db.Take(&category, "id = ?", id).Error

	return category, err
}
