package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"

	"todolib/db"
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
