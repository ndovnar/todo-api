package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"todolib/mongodb"

	"todoauth/internal/model"
)

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *userRepository {
	collection := db.Collection("users")

	return &userRepository{
		collection: collection,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	currentTime := time.Now()
	user.Dates = model.Dates{
		Created:  &currentTime,
		Modified: &currentTime,
	}

	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, mongodb.MongoErrorToDBError(err)
	}

	newID, err := mongodb.InsertedIDToHex(result.InsertedID)
	if err != nil {
		return nil, err
	}

	user.ID = newID

	return user, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	objectID, err := mongodb.IDHexToObjectID(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id":     objectID,
		"deleted": false,
	}

	return r.getUser(ctx, filter)
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	filter := bson.M{
		"email":   email,
		"deleted": false,
	}

	return r.getUser(ctx, filter)
}

func (r *userRepository) getUser(ctx context.Context, filter bson.M) (*model.User, error) {
	res := r.collection.FindOne(ctx, filter)
	user := &model.User{}

	err := res.Decode(user)
	if err != nil {
		return nil, fmt.Errorf("failed to decode user: %w", err)
	}

	return user, nil
}
