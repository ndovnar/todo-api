package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/sync/errgroup"

	"todolib/db"
	"todolib/mongodb"

	"todo/internal/modeldb"
	"todo/internal/repository"
)

type todoRepository struct {
	collection *mongo.Collection
}

func NewTodoRepository(db *mongo.Database) *todoRepository {
	collection := db.Collection("todo")

	return &todoRepository{
		collection: collection,
	}
}

func (r *todoRepository) GetTodos(ctx context.Context, arg *repository.GetTodosParams) ([]*modeldb.Todo, int64, error) {
	errGroup, gCtx := errgroup.WithContext(ctx)
	filter := bson.M{
		"userId":  arg.UserID,
		"deleted": false,
	}
	options := &options.FindOptions{
		Limit: &arg.Limit,
		Skip:  &arg.Offset,
	}

	todos := []*modeldb.Todo{}
	errGroup.Go(func() error {
		cursor, err := r.collection.Find(ctx, filter, options)
		if err != nil {
			return fmt.Errorf("failed to get todos, %w", err)
		}

		if err := cursor.All(gCtx, &todos); err != nil {
			return fmt.Errorf("failed to decode todos, %w", err)
		}

		return nil
	})

	var totalCount int64
	errGroup.Go(func() error {
		count, err := r.collection.CountDocuments(ctx, filter)
		if err != nil {
			return fmt.Errorf("failed to count todos, %w", err)
		}

		totalCount = count
		return nil
	})

	if err := errGroup.Wait(); err != nil {
		return nil, 0, err
	}

	return todos, totalCount, nil
}

func (r *todoRepository) GetTodo(ctx context.Context, arg *repository.GetTodoParams) (*modeldb.Todo, error) {
	return r.getTodoByID(ctx, arg.ID, arg.UserID)
}

func (r todoRepository) UpdateTodo(ctx context.Context, arg *repository.UpdateTodoParams) (*modeldb.Todo, error) {
	objectID, err := mongodb.IDHexToObjectID(arg.ID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	filter := bson.M{
		"_id":     objectID,
		"deleted": false,
		"userId":  arg.UserID,
	}
	currentTime := time.Now()
	update := bson.M{
		"$set": bson.M{
			"title":          arg.Title,
			"description":    arg.Description,
			"completed":      arg.IsCompleted,
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

	return r.getTodoByID(ctx, arg.ID, arg.UserID)
}

func (r *todoRepository) CreateTodo(ctx context.Context, arg *repository.CreateTodoParams) (*modeldb.Todo, error) {
	currentTime := time.Now()
	todo := &modeldb.Todo{
		Title:       arg.Title,
		Description: arg.Description,
		UserID:      arg.UserID,
		Dates: &modeldb.Dates{
			Created:  &currentTime,
			Modified: &currentTime,
		},
	}

	result, err := r.collection.InsertOne(ctx, todo)
	if err != nil {
		return nil, mongodb.MongoErrorToDBError(err)
	}

	id, err := mongodb.InsertedIDToHex(result.InsertedID)
	if err != nil {
		return nil, err
	}

	return r.getTodoByID(ctx, id, arg.UserID)
}

func (r *todoRepository) DeleteTodo(ctx context.Context, arg *repository.DeleteTodoParams) error {
	objectID, err := mongodb.IDHexToObjectID(arg.ID)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id":     objectID,
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

func (r *todoRepository) getTodoByID(ctx context.Context, id, userId string) (*modeldb.Todo, error) {
	objectID, err := mongodb.IDHexToObjectID(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id":     objectID,
		"userId":  userId,
		"deleted": false,
	}

	result := r.collection.FindOne(ctx, filter)
	todo := &modeldb.Todo{}
	err = result.Decode(todo)
	if err != nil {
		return nil, mongodb.MongoErrorToDBError(err)
	}

	return todo, nil
}
