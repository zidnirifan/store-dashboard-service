package handler

import (
	"net/http"
	"store-dashboard-service/model"
	"store-dashboard-service/util/exception"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	validationError, ok := err.(validator.ValidationErrors)
	if ok {
		return c.Status(http.StatusBadRequest).JSON(model.Response{
			Message: getValidationErrorMsg(validationError[0]),
			Success: false,
		})
	}

	re, ok := err.(*exception.CustomError)
	if ok {
		return c.Status(re.StatusCode).JSON(model.Response{
			Message: re.Error(),
			Success: false,
		})
	}

	return c.Status(http.StatusInternalServerError).JSON(model.Response{
		Message: err.Error(),
		Success: false,
	})
}

func getValidationErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "email":
		return fe.Field() + " must be a valid email address"
	case "min":
		return fe.Field() + " must be at least " + fe.Param() + " characters long"
	case "max":
		return fe.Field() + " cannot be longer than " + fe.Param() + " characters"
	case "gte":
		return fe.Field() + " must be greater than or equal to " + fe.Param()
	case "lte":
		return fe.Field() + " must be less than or equal to " + fe.Param()
	}
	return fe.Field() + " is invalid"
}
