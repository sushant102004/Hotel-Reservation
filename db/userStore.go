package db

import (
	"context"

	"github.com/sushant102004/Hotel-Reservation-System/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// This interface is designed so that we can work with different databases at same time.

type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
	col    *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	col := client.Database(DBNAME).Collection("users")

	return &MongoUserStore{
		client: client,
		col:    col,
	}
}

func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	var user *types.User

	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if err := s.col.FindOne(ctx, bson.M{"_id": oID}).Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}
