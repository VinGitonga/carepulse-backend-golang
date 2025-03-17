package main

import (
	"CarepluseBackend/database"
	"CarepluseBackend/router"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber", AppName: "Carepulse"})

	// Connect to db
	database.ConnectDB()

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":4610"))
}
