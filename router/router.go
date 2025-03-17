package router

import (
	"CarepluseBackend/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	api.Get("/", handler.Hello)

	// User
	user := api.Group("/users")
	user.Get("/", handler.GetUsers)
	user.Get("/profile/:id", handler.GetUser)
	user.Post("/admin/create", handler.CreateAdminUser)
	user.Post("/patient/create", handler.CreateNewPatient)
}
