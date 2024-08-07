package api

import (
	"store-dashboard-service/api/route"

	"github.com/gofiber/fiber/v2"
)

func NewServer(userRoute *route.UserRoute) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "store-dashboard-service",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	userRouter := app.Group("/users")
	userRoute.Init(userRouter)

	return app
}
