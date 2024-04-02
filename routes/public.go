package routes

import (
	"rinto11804/activity_point_backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupPublicRoutes(app *fiber.App) {
	api := app.Group("/api")
	auth := api.Group("/auth")
	auth.Post("/signup/:role", controllers.SignUp)
	auth.Post("/signin/:role", controllers.SignIn)
}
