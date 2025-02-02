package mongodb

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertedIDToObjectID(insertedID any) (*primitive.ObjectID, error) {
	objectID, ok := insertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("id of the inserted document %q is not an object id", insertedID)
	}

	return &objectID, nil
}

func InsertedIDToHex(insertedID any) (string, error) {
	objectID, err := InsertedIDToObjectID(insertedID)
	if err != nil {
		return "", err
	}

	return objectID.Hex(), nil
}

func IDHexToObjectID(id string) (*primitive.ObjectID, error) {
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("failed to convert hex to ObjectID %w", err)
	}

	return &primitiveID, nil
}
