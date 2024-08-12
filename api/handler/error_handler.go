package handler

import (
	"net/http"
	"store-dashboard-service/model"
	"store-dashboard-service/util/exception"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	re, ok := err.(*exception.CustomError)
	if ok {
		c.Status(re.StatusCode)
		return c.JSON(model.Response{
			Message: re.Error(),
			Success: false,
		})
	}

	c.Status(http.StatusInternalServerError)
	return c.JSON(model.Response{
		Message: err.Error(),
		Success: false,
	})
}
