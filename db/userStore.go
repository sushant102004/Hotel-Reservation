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
	GetUsers(context.Context) ([]*types.User, error)
	InsertUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(context.Context, bson.M, bson.M) error
	GetUserByEmail(context.Context, string) (*types.User, error)
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

func (s *MongoUserStore) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	var user types.User
	if err := s.col.FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	var users []*types.User

	cur, err := s.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cur.All(ctx, &users); err != nil {
		return []*types.User{}, nil
	}

	return users, nil
}

func (s *MongoUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := s.col.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	user.ID = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (s *MongoUserStore) DeleteUser(ctx context.Context, userID string) error {
	oID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return err
	}

	_, err = s.col.DeleteOne(ctx, bson.M{"_id": oID})
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoUserStore) UpdateUser(ctx context.Context, filter, update bson.M) error {
	_, err := s.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
