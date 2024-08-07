package api

import "github.com/gofiber/fiber/v2"

func NewServer() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "store-dashboard-service",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	return app
}
