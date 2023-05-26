package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurbekchymbaev/sanctions-go-app/database"
)

func main() {
	database.DbConnect()

	app := fiber.New()
	setupRoutes(app)

	app.Listen(":8080")
}
