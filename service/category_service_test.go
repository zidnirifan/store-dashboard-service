package service

import (
	"errors"
	"store-dashboard-service/model"
	"store-dashboard-service/repository/mocks"
	"store-dashboard-service/util/exception"
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

var categoryFromDB = model.Category{
	ID:   1,
	Name: "test",
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

func TestGetCategories(t *testing.T) {
	t.Run("it should success", func(t *testing.T) {
		categoryService, categoryRepository := newCategoryServiceWithMock()
		categoryRepository.On("GetAll", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*[]model.Category)
			*arg = append(*arg, categoryFromDB)
		})

		categories, err := categoryService.GetCategories()

		assert.Nil(t, err)
		assert.NotEmpty(t, categories)
		categoryRepository.AssertExpectations(t)
	})

	t.Run("it should return error", func(t *testing.T) {
		categoryService, categoryRepository := newCategoryServiceWithMock()
		categoryRepository.On("GetAll", mock.Anything).Return(errors.New("error"))

		_, err := categoryService.GetCategories()

		assert.NotNil(t, err)
		categoryRepository.AssertExpectations(t)
	})
}

func TestGetCategoryById(t *testing.T) {
	t.Run("it should success", func(t *testing.T) {
		categoryService, categoryRepository := newCategoryServiceWithMock()
		categoryRepository.On("GetById", mock.Anything).Return(categoryFromDB, nil)

		category, err := categoryService.GetCategoryById(1)

		assert.Nil(t, err)
		assert.NotEmpty(t, category)
		categoryRepository.AssertExpectations(t)
	})

	t.Run("it should return 404 error", func(t *testing.T) {
		categoryService, categoryRepository := newCategoryServiceWithMock()
		categoryRepository.On("GetById", mock.Anything).Return(model.Category{}, errors.New("error"))

		category, err := categoryService.GetCategoryById(11)

		assert.NotNil(t, err)
		assert.Empty(t, category)
		assert.Equal(t, 404, err.(*exception.CustomError).StatusCode)
		categoryRepository.AssertExpectations(t)
	})
}
