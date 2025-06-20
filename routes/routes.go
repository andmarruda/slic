package routes

import (
	"github.com/gofiber/fiber/v2"
	"slic/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.group("/api/v1")

	api.Post("/aws-s3/upload", handlers.Upload)
}
