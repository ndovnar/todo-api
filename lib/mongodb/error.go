package mongodb

import (
	"lib/db"

	"go.mongodb.org/mongo-driver/mongo"
)

func MongoErrorToDBError(err error) error {
	if err == mongo.ErrNoDocuments {
		return db.ErrNotFound
	}

	if mongo.IsDuplicateKeyError(err) {
		return db.ErrDuplicateKey
	}

	return db.ErrUnknown
}
