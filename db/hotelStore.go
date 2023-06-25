package db

import (
	"context"

	"github.com/sushant102004/Hotel-Reservation-System/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelStore interface {
}

type MongoHotelStore struct {
	client *mongo.Client
	col    *mongo.Collection
}

func NewHotelStore(client *mongo.Client) *MongoHotelStore {
	col := client.Database(DBNAME).Collection("hotels")

	return &MongoHotelStore{
		client: client,
		col:    col,
	}
}

func (s *MongoHotelStore) InsertHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	resp, err := s.col.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}

	hotel.ID = resp.InsertedID.(primitive.ObjectID)
	return hotel, nil
}
