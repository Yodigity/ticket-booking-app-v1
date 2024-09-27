package main

import (
	"github.com/Yodigity/ticket-booking-app-v1/handlers"
	"github.com/Yodigity/ticket-booking-app-v1/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New((fiber.Config{
		AppName:      "Ticketbooking",
		ServerHeader: "Fiber",
	}))

	// Repositories
	eventRepository := repositories.NewEventRepository(nil)

	// Routing
	server := app.Group("/api")

	// Handlers

	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(":3000")
}
