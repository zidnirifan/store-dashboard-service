package handler

import (
	"store-dashboard-service/model"
	"store-dashboard-service/service"
	"strconv"

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

func (ch *CategoryHandler) GetCategoryById(c *fiber.Ctx) error {
	idString := c.Params("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(model.Response{
			Success: false,
			Message: "category not found",
		})
	}

	category, err := ch.service.GetCategoryById(id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(model.Response{
		Success: true,
		Message: "success get category",
		Data:    category,
	})
}

func (ch *CategoryHandler) UpdateCategory(c *fiber.Ctx) error {
	idString := c.Params("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(model.Response{
			Success: false,
			Message: "category not found",
		})
	}
	body := &model.UpdateCategoryRequest{}
	err = c.BodyParser(body)
	if err != nil {
		return err
	}

	category, err := ch.service.UpdateCategory(id, body)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(model.Response{
		Success: true,
		Message: "update category successfully",
		Data:    category,
	})
}

func (ch *CategoryHandler) DeleteCategory(c *fiber.Ctx) error {
	idString := c.Params("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(model.Response{
			Success: false,
			Message: "category not found",
		})
	}

	err = ch.service.DeleteCategoryById(id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(model.Response{
		Success: true,
		Message: "delete category successfully",
	})
}
