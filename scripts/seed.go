package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sushant102004/Hotel-Reservation-System/db"
	"github.com/sushant102004/Hotel-Reservation-System/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	const dbURI = "mongodb://localhost:27017"
	ctx := context.Background()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbURI))
	if err != nil {
		panic(err)
	}

	client.Database(db.DBNAME).Drop(ctx)

	hotelStore := db.NewHotelStore(client)
	roomStore := db.NewRoomStore(client, hotelStore)

	hotel := types.Hotel{
		Name:     "Daddy's Hotel",
		Location: "California",
		Rooms:    []primitive.ObjectID{},
	}

	rooms := []types.Room{
		{
			Type:      types.DoubleRoom,
			BasePrice: 99,
		},
		{
			Type:      types.SingleRoom,
			BasePrice: 59,
		},
		{
			Type:      types.LuxryRoom,
			BasePrice: 159,
		},
	}

	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, room := range rooms {
		room.HotelID = insertedHotel.ID
		insertedRoom, err := roomStore.InsertRoom(ctx, &room)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(insertedRoom)
	}

	fmt.Println(insertedHotel)

}
