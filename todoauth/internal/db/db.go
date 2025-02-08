package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateIndexes(ctx context.Context, database *mongo.Database) error {
	usersCollection := database.Collection("users")
	_, err := usersCollection.Indexes().CreateMany(
		ctx,
		[]mongo.IndexModel{
			{Keys: bson.M{"email": 1}, Options: options.Index().SetUnique(true)},
		},
	)

	if err != nil {
		return err
	}

	return err
}
