package main

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Connect to database
	db, error := ConnectDatabase()

	if error != nil {
		fmt.Println(error, db)
	}

	// Initialize Fiber
	app := fiber.New()

	// Routes Activate
	studentRoute(app, db)

	// Start server
	if error := app.Listen(":1112"); error != nil {
		fmt.Println(error)
	}
}
