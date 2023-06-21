package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	app := fiber.New()
	const dbURI = "mongodb://localhost:27017"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbURI))

	if err != nil {
		panic(err)
	}

	app.Listen(":5000")
}
