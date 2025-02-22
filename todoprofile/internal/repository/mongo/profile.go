package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"todolib/db"
	"todolib/mongodb"

	"todoprofile/internal/modeldb"
	"todoprofile/internal/repository"
)

type profileRepository struct {
	collection *mongo.Collection
}

func NewProfileRepository(db *mongo.Database) *profileRepository {
	collection := db.Collection("profiles")

	return &profileRepository{
		collection: collection,
	}
}

func (r *profileRepository) GetProfile(ctx context.Context, arg *repository.GetProfileParams) (*modeldb.Profile, error) {
	return r.getProfileByUserID(ctx, arg.UserID)
}

func (r *profileRepository) CreateProfile(ctx context.Context, arg *repository.CreateProfileParams) (*modeldb.Profile, error) {
	currentTime := time.Now()
	profile := &modeldb.Profile{
		UserID: arg.UserID,
		Dates: &modeldb.Dates{
			Created:  &currentTime,
			Modified: &currentTime,
		},
	}

	_, err := r.collection.InsertOne(ctx, profile)
	if err != nil {
		return nil, mongodb.MongoErrorToDBError(err)
	}

	return r.getProfileByUserID(ctx, arg.UserID)
}

func (r profileRepository) UpdateProfile(ctx context.Context, arg *repository.UpdateProfileParams) (*modeldb.Profile, error) {
	filter := bson.M{
		"deleted": false,
		"userId":  arg.UserID,
	}
	currentTime := time.Now()
	update := bson.M{
		"$set": bson.M{
			"firstName":      arg.FirstName,
			"lastName":       arg.LastName,
			"dates.modified": &currentTime,
		},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, db.ErrNotFound
	}

	return r.getProfileByUserID(ctx, arg.UserID)
}

func (r *profileRepository) DeleteProfile(ctx context.Context, arg *repository.DeleteProfileParams) error {
	filter := bson.M{
		"deleted": false,
		"userId":  arg.UserID,
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

func (r *profileRepository) getProfileByUserID(ctx context.Context, userId string) (*modeldb.Profile, error) {
	filter := bson.M{
		"userId":  userId,
		"deleted": false,
	}

	result := r.collection.FindOne(ctx, filter)
	profile := &modeldb.Profile{}
	err := result.Decode(profile)
	if err != nil {
		return nil, mongodb.MongoErrorToDBError(err)
	}

	return profile, nil
}
