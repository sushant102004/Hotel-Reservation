package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sushant102004/Hotel-Reservation-System/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	user := types.User{
		FirstName: "Sushant",
		LastName:  "Dhiman",
	}
	return c.JSON(user)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("Sushant")
}
