package api

import (
	"store-dashboard-service/api/handler"
	"store-dashboard-service/api/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func NewServer(userRoute *route.UserRoute) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      "store-dashboard-service",
		ErrorHandler: handler.ErrorHandler,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Get("/swagger.yaml", func(c *fiber.Ctx) error {
		return c.SendFile("./swagger.yaml")
	})
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL: "/swagger.yaml",
	}))

	userRouter := app.Group("/users")
	userRoute.Init(userRouter)

	return app
}
