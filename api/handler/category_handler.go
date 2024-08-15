package handler

import (
	"store-dashboard-service/model"
	"store-dashboard-service/service"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (ch *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	body := &model.CreateCategoryRequest{}
	err := c.BodyParser(body)
	if err != nil {
		return err
	}

	category, err := ch.service.CreateCategory(body)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(model.Response{
		Success: true,
		Message: "category created successfully",
		Data:    category,
	})
}

func (ch *CategoryHandler) GetCategories(c *fiber.Ctx) error {
	categories, err := ch.service.GetCategories()
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(model.Response{
		Success: true,
		Message: "success get categories",
		Data:    categories,
	})
}
