package handler

import (
	"store-dashboard-service/model"
	"store-dashboard-service/service"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (ch *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	body := &model.CreateProductRequest{}
	err := c.BodyParser(body)
	if err != nil {
		return err
	}
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["images"]
	body.Images = files

	product, err := ch.service.CreateProduct(c.Context(), body)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(model.Response{
		Success: true,
		Message: "product created successfully",
		Data:    product,
	})
}
