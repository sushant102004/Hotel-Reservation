package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sushant102004/Hotel-Reservation-System/api"
)

func main() {
	app := fiber.New()

	app.Get("/foo", handleFoo)

	apiV1 := app.Group("/api/v1/")

	apiV1.Get("/user", api.HandleGetUsers)
	apiV1.Get("/user/:id", api.HandleGetUser)

	app.Listen(":5000")
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "Server Working"})
}
