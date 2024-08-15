package service

import (
	"errors"
	"store-dashboard-service/model"
	"store-dashboard-service/repository/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func newCategoryServiceWithMock() (*CategoryService, *mocks.CategoryRepository) {
	categoryRepository := new(mocks.CategoryRepository)
	categoryService := NewCategoryService(categoryRepository, validator.New())
	return categoryService, categoryRepository
}

func TestCreateCategory(t *testing.T) {
	t.Run("it should success", func(t *testing.T) {
		payload := model.CreateCategoryRequest{
			Name: "test",
		}

		categoryService, categoryRepository := newCategoryServiceWithMock()
		categoryRepository.On("Create", mock.Anything).Return(nil)

		category, err := categoryService.CreateCategory(&payload)

		assert.Nil(t, err)
		assert.NotEmpty(t, category)
		categoryRepository.AssertExpectations(t)
	})

	t.Run("it should return error", func(t *testing.T) {
		payload := model.CreateCategoryRequest{
			Name: "test",
		}

		categoryService, categoryRepository := newCategoryServiceWithMock()
		categoryRepository.On("Create", mock.Anything).Return(errors.New("error"))

		_, err := categoryService.CreateCategory(&payload)

		assert.NotNil(t, err)
		categoryRepository.AssertExpectations(t)
	})
}
