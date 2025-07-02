package middleware

import (
	"strings"

	"github.com/yeawyow/gateway/util"

	"github.com/gofiber/fiber/v2"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing token"})
		}

		token := strings.TrimPrefix(auth, "Bearer ")

		claims, err := util.ParseJWT(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
		}

		c.Locals("user_id", claims["user_id"])
		return c.Next()
	}
}
