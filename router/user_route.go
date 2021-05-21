package routes

import (
	"gotbotpoc/controllers"

	"github.com/gofiber/fiber"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {

	// Middleware
	api := app.Group("/api")

	// routes
	api.Get("/getUserList", controllers.GetUserList)
	api.Post("/addUpdateUser", controllers.AddUpdateUser)
	api.Delete("/deleteUser/:id", controllers.DeleteUser)
	api.Post("/login", controllers.Login)

}
