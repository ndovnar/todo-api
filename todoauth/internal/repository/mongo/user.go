package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"todolib/mongodb"

	"todoauth/internal/modeldb"
	"todoauth/internal/repository"
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

func (r *userRepository) CreateUser(ctx context.Context, arg *repository.CreateUserParams) (*modeldb.User, error) {
	currentTime := time.Now()

	user := &modeldb.User{
		Email:    arg.Email,
		Password: arg.Password,
		Dates: &modeldb.Dates{
			Created:  &currentTime,
			Modified: &currentTime,
		},
	}

	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, mongodb.MongoErrorToDBError(err)
	}

	objectID, err := mongodb.InsertedIDToObjectID(result.InsertedID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id":     objectID,
		"deleted": false,
	}

	return r.getUser(ctx, filter)
}

func (r *userRepository) GetUserByID(ctx context.Context, id string) (*modeldb.User, error) {
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

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*modeldb.User, error) {
	filter := bson.M{
		"email":   email,
		"deleted": false,
	}

	return r.getUser(ctx, filter)
}

func (r *userRepository) getUser(ctx context.Context, filter bson.M) (*modeldb.User, error) {
	res := r.collection.FindOne(ctx, filter)
	user := &modeldb.User{}

	err := res.Decode(user)
	if err != nil {
		return nil, fmt.Errorf("failed to decode user: %w", err)
	}

	return user, nil
}
