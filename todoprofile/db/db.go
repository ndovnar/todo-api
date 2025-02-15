package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateIndexes(ctx context.Context, database *mongo.Database) error {
	fmt.Printf("\"CreateIndexes\": %v\n", "CreateIndexes")
	usersCollection := database.Collection("profiles")
	_, err := usersCollection.Indexes().CreateMany(
		ctx,
		[]mongo.IndexModel{
			{Keys: bson.M{"userId": 1}, Options: options.Index().SetUnique(true)},
		},
	)

	if err != nil {
		return err
	}

	return err
}
