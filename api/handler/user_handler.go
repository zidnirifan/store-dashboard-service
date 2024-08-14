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
		Success: true,
		Data:    token,
	})
}

func (handler *UserHandler) Register(c *fiber.Ctx) error {
	body := &model.RegisterRequest{}
	err := c.BodyParser(body)
	if err != nil {
		return err
	}

	res, err := handler.userService.Register(body)
	if err != nil {
		return err
	}

	return c.JSON(model.Response{
		Message: "login success",
		Success: true,
		Data:    res,
	})
}

func (handler *UserHandler) VerifyUser(c *fiber.Ctx) error {
	userId := c.Params("userId")

	err := handler.userService.VerifyUser(userId)
	if err != nil {
		return err
	}

	return c.JSON(model.Response{
		Message: "user verified successfully",
		Success: true,
	})
}
