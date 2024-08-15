package api

import (
	"store-dashboard-service/api/handler"
	"store-dashboard-service/api/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

type Route struct {
	UserRoute     *route.UserRoute
	CategoryRoute *route.CategoryRoute
}

func NewServer(route *Route) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      "store-dashboard-service",
		ErrorHandler: handler.ErrorHandler,
	})

	app.Use(logger.New())
	app.Use(healthcheck.New())

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
	route.UserRoute.Init(userRouter)

	categoryRouter := app.Group("/categories")
	route.CategoryRoute.Init(categoryRouter)

	return app
}
