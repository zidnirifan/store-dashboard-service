package route

import (
	"store-dashboard-service/api/handler"

	"github.com/gofiber/fiber/v2"
)

type CategoryRoute struct {
	handler *handler.CategoryHandler
}

func NewCategoryRoute(handler *handler.CategoryHandler) *CategoryRoute {
	return &CategoryRoute{handler: handler}
}

func (categoryRoute *CategoryRoute) Init(router fiber.Router) {
	router.Post("/", categoryRoute.handler.CreateCategory)
	router.Get("/", categoryRoute.handler.GetCategories)
	router.Get("/:id", categoryRoute.handler.GetCategoryById)
	router.Put("/:id", categoryRoute.handler.UpdateCategory)
	router.Delete("/:id", categoryRoute.handler.DeleteCategory)
}
