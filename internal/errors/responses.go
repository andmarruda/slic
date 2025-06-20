package errors

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorJSON(c *fiber.Ctx, status int, msg string, err error) error {
	return c.Status(status).JSON(fiber.Map{
		"message": msg,
		"status":  "error",
		"error":   err.Error(),
	})
}
