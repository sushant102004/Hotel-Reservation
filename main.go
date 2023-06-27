package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/sushant102004/Hotel-Reservation-System/api"
	"github.com/sushant102004/Hotel-Reservation-System/db"
	"github.com/sushant102004/Hotel-Reservation-System/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	app := fiber.New(config)
	const dbURI = "mongodb://localhost:27017"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbURI))
	if err != nil {
		panic(err)
	}

	var (
		userHandler  = api.NewUserHandler(db.NewMongoUserStore(client))
		hotelStore   = db.NewMongoHotelStore(client)
		roomStore    = db.NewMongoRoomStore(client, hotelStore)
		hotelHandler = api.NewHotelHandler(hotelStore, roomStore)
	)

	apiVOne := app.Group("/api/v1", middleware.Authenticate)

	apiVOne.Get("/user", userHandler.HandleGetUsers)
	apiVOne.Get("/user/:id", userHandler.HandleGetUser)
	apiVOne.Post("/user", userHandler.HandleInsertUser)
	apiVOne.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiVOne.Put("/user/:id", userHandler.HandleUpdateUser)

	apiVOne.Get("/hotel", hotelHandler.HandleGetHotels)
	apiVOne.Get("/hotel/:id/rooms", hotelHandler.HandleGetRooms)
	apiVOne.Get("/hotel/:id", hotelHandler.HandleGetHotel)

	app.Listen(":5000")
}
