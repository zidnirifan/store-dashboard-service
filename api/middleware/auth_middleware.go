package middleware

import (
	"store-dashboard-service/config"
	"store-dashboard-service/model"
	"store-dashboard-service/util"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func SuperAdminAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
				Success: false,
				Message: "authorization is required",
			})
		}

		tokenString := authHeader[len("Bearer "):]
		user := &model.PayloadAccessToken{}
		token, err := jwt.ParseWithClaims(tokenString, user, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.GetConfig().AccessTokenKey), nil
		})
		if !token.Valid || err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
				Success: false,
				Message: "invalid token",
			})
		}
		if user.Role != util.CommonConst.Roles.SuperAdmin {
			return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
				Success: false,
				Message: "access forbidden for this user",
			})
		}

		c.Locals("user", user)
		return c.Next()
	}
}
