package route

import (
	"store-dashboard-service/api/handler"

	"github.com/gofiber/fiber/v2"
)

type ProductRoute struct {
	handler *handler.ProductHandler
}

func NewProductRoute(handler *handler.ProductHandler) *ProductRoute {
	return &ProductRoute{handler: handler}
}

func (categoryRoute *ProductRoute) Init(router fiber.Router) {
	router.Post("/", categoryRoute.handler.CreateProduct)
}
