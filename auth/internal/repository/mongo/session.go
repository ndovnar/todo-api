package mongo

import (
	"auth/internal/model"
	"context"
	"fmt"
	"time"

	"lib/db"
	"lib/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func (r *sessionRepository) GetSessionByID(ctx context.Context, id string) (*model.Session, error) {
	objectID, err := mongodb.IDHexToObjectID(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id":     objectID,
		"deleted": false,
	}

	result := r.collection.FindOne(ctx, filter)
	session := &model.Session{}
	err = result.Decode(session)

	if err != nil {
		return nil, fmt.Errorf("failed to decode session: %w", err)
	}

	return session, nil
}

func (r *sessionRepository) CreateSession(ctx context.Context, session *model.Session) (*model.Session, error) {
	currentTime := time.Now()
	session.Dates = model.Dates{
		Created:  &currentTime,
		Modified: &currentTime,
	}

	result, err := r.collection.InsertOne(ctx, session)
	if err != nil {
		return nil, mongodb.MongoErrorToDBError(err)
	}

	id, err := mongodb.InsertedIDToHex(result.InsertedID)
	if err != nil {
		return nil, err
	}

	session.ID = id
	return session, nil
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
