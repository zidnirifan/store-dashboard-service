package service

import (
	"store-dashboard-service/model"
	"store-dashboard-service/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryService struct {
	repository repository.CategoryRepository
	validate   *validator.Validate
}

func NewCategoryService(repository repository.CategoryRepository, validate *validator.Validate) *CategoryService {
	return &CategoryService{repository: repository, validate: validate}
}

func (c *CategoryService) CreateCategory(payload *model.CreateCategoryRequest) (model.Category, error) {
	var category model.Category

	err := c.validate.Struct(payload)
	if err != nil {
		return category, err
	}

	category = model.Category{
		Name: payload.Name,
	}
	err = c.repository.Create(&category)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (c *CategoryService) GetCategories() ([]model.Category, error) {
	var categories []model.Category

	err := c.repository.GetAll(&categories)
	if err != nil {
		return categories, err
	}

	return categories, nil
}
