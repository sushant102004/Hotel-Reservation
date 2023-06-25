package db

import (
	"context"

	"github.com/sushant102004/Hotel-Reservation-System/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomStore interface {
}

type MongoRoomStore struct {
	client *mongo.Client
	col    *mongo.Collection
}

func NewRoomStore(client *mongo.Client) *MongoRoomStore {
	col := client.Database(DBNAME).Collection("rooms")

	return &MongoRoomStore{
		client: client,
		col:    col,
	}
}

func (s *MongoRoomStore) InsertRoom(ctx context.Context, room *types.Room) (*types.Room, error) {
	resp, err := s.col.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}

	room.ID = resp.InsertedID.(primitive.ObjectID)
	return room, nil
}
