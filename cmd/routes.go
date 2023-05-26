package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurbekchymbaev/sanctions-go-app/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/update", handlers.Update)
	app.Get("/state", handlers.State)
	app.Get("/get_names", handlers.Getnames)
}
