package db

import (
	"context"

	"github.com/sushant102004/Hotel-Reservation-System/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomStore interface {
	InsertRoom(context.Context, *types.Room) (*types.Room, error)
}

type MongoRoomStore struct {
	client     *mongo.Client
	col        *mongo.Collection
	hotelStore HotelStore
}

func NewRoomStore(client *mongo.Client, hotelStore HotelStore) *MongoRoomStore {
	col := client.Database(DBNAME).Collection("rooms")

	return &MongoRoomStore{
		client:     client,
		col:        col,
		hotelStore: hotelStore,
	}
}

func (s *MongoRoomStore) InsertRoom(ctx context.Context, room *types.Room) (*types.Room, error) {
	resp, err := s.col.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}

	room.ID = resp.InsertedID.(primitive.ObjectID)

	// Adding this room to hotel document.
	filter := bson.M{"_id": room.HotelID}
	update := bson.M{"$push": bson.M{"rooms": room.ID}}

	if err := s.hotelStore.Update(context.Background(), filter, update); err != nil {
		return nil, err
	}

	return room, nil
}
