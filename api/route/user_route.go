package route

import (
	"store-dashboard-service/api/handler"
	"store-dashboard-service/api/middleware"

	"github.com/gofiber/fiber/v2"
)

type UserRoute struct {
	handler *handler.UserHandler
}

func NewUserRoute(handler *handler.UserHandler) *UserRoute {
	return &UserRoute{handler: handler}
}

func (userRoute *UserRoute) Init(router fiber.Router) {
	router.Post("/login", userRoute.handler.Login)
	router.Post("/register", userRoute.handler.Register)
	router.Put("/verify/:userId", middleware.SuperAdminAuth(), userRoute.handler.VerifyUser)
}
