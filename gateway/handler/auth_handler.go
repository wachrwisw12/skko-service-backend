package handler

import "github.com/gofiber/fiber/v2"

func LineHandler(c *fiber.Ctx) error {
	Code :=c.Query("code")
	State :=c.Query("state")
	return c.JSON(fiber.Map{
		"cdoe":Code,
		"state":State,
	})
}