package handler

import (
	"store-dashboard-service/model"
	"store-dashboard-service/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (handler *UserHandler) Login(c *fiber.Ctx) error {
	body := &model.LoginRequest{}
	err := c.BodyParser(body)
	if err != nil {
		return err
	}

	token, err := handler.userService.Login(body)
	if err != nil {
		return err
	}

	return c.JSON(model.Response{
		Message: "login success",
		Data:    token,
	})
}
