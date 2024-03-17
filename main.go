package main

import (
	"Simple-REST-ful-Fiber/database"
	"Simple-REST-ful-Fiber/router"

	"github.com/gofiber/fiber/v2"
)
func main() {
	// Start a new fiber app
	app := fiber.New()

	// Connect to the Database
	db := database.ConnectDB()

	// Setup the router
	router.InitRouter(app, db)

	// Listen on port 8000
	app.Listen(":8000")
}