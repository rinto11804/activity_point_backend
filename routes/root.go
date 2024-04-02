package routes

import (
	"rinto11804/activity_point_backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/healthcheck", controllers.HealthCheck)
}
