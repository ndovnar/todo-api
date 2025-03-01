package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"todolib/db"
	"todolib/mongodb"

	"todoauth/internal/modeldb"
)

type sessionRepository struct {
	collection *mongo.Collection
}

func NewSessionRepository(db *mongo.Database) *sessionRepository {
	collection := db.Collection("sessions")

	return &sessionRepository{
		collection: collection,
	}
}

func (r *sessionRepository) GetSessionByID(ctx context.Context, id string) (*modeldb.Session, error) {
	return r.getSessionByID(ctx, id)
}

func (r *sessionRepository) CreateSession(ctx context.Context, userID string) (*modeldb.Session, error) {
	currentTime := time.Now()
	session := &modeldb.Session{
		UserID: userID,
		Dates: &modeldb.Dates{
			Created:  &currentTime,
			Modified: &currentTime,
		},
	}

	result, err := r.collection.InsertOne(ctx, session)
	if err != nil {
		return nil, mongodb.MongoErrorToDBError(err)
	}

	id, err := mongodb.InsertedIDToHex(result.InsertedID)
	if err != nil {
		return nil, err
	}

	return r.getSessionByID(ctx, id)
}

func (r *sessionRepository) DeleteSession(ctx context.Context, id string) error {
	objectID, err := mongodb.IDHexToObjectID(id)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id":     objectID,
		"deleted": false,
	}

	update := bson.M{
		"$set": bson.M{
			"deleted":       true,
			"dates.deleted": time.Now(),
		},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return mongodb.MongoErrorToDBError(err)
	}

	if result.MatchedCount == 0 {
		return db.ErrNotFound
	}

	return nil
}

func (r *sessionRepository) getSessionByID(ctx context.Context, id string) (*modeldb.Session, error) {
	objectID, err := mongodb.IDHexToObjectID(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id":     objectID,
		"deleted": false,
	}

	result := r.collection.FindOne(ctx, filter)
	session := &modeldb.Session{}

	err = result.Decode(session)
	if err != nil {
		return nil, mongodb.MongoErrorToDBError(err)
	}

	return session, nil
}
