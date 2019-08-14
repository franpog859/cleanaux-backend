package database

import (
	"context"
	"fmt"

	"github.com/franpog859/cleanaux-backend/auth-service/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoDBServiceURI = "mongodb://mongo-database-internal:27017"
	databaseName      = "users"
	collectionName    = "basicauth"
)

// Client interface
type Client interface {
	GetAuthorizedUsers(string, string) ([]model.User, error)
}

type client struct {
	collection *mongo.Collection
}

// NewClient provides Client interface
func NewClient() (Client, error) {
	clientOptions := options.Client().ApplyURI(mongoDBServiceURI)
	mongoDBClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect MongoDB client: %v", err)
	}

	collection := mongoDBClient.Database(databaseName).Collection(collectionName)

	return &client{
		collection: collection,
	}, nil
}

func (c *client) GetAuthorizedUsers(username, password string) ([]model.User, error) {
	filter := bson.M{
		"username": username,
		"password": password,
	}

	cursor, err := c.collection.Find(context.TODO(), filter, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to find users in collection: %v", err)
	}
	defer cursor.Close(context.TODO())

	users, err := getUsersFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func getUsersFromCursor(cursor *mongo.Cursor) ([]model.User, error) {
	var results []model.User

	for cursor.Next(context.TODO()) {
		var user model.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, fmt.Errorf("failed to decode user from cursor: %v", err)
		}

		results = append(results, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return results, nil
}
