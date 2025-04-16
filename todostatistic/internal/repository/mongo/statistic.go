package mongo

import (
	"context"
	"todolib/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"todostatistic/internal/modeldb"
	"todostatistic/internal/repository"
)

type statisticRepository struct {
	collection *mongo.Collection
}

func NewStatisticRepository(db *mongo.Database) *statisticRepository {
	collection := db.Collection("statistic")
	return &statisticRepository{
		collection: collection,
	}
}

func (r *statisticRepository) GetUserTodoStatistic(ctx context.Context, arg *repository.GetUserTodoStatisticParams) (*modeldb.UserTodoStatistic, error) {
	month, err := r.getUserTodoStatisticMonthly(ctx, arg.UserID, arg.Month)
	if err != nil {
		return nil, mongodb.MongoErrorToDBError(err)
	}

	year, err := r.getUserTodoStatisticYearly(ctx, arg.UserID, arg.Year)
	if err != nil {
		return nil, mongodb.MongoErrorToDBError(err)
	}

	total, err := r.getStatisticTotal(ctx, arg.UserID)
	if err != nil {
		return nil, mongodb.MongoErrorToDBError(err)
	}

	return &modeldb.UserTodoStatistic{
		Total: total,
		Year:  year,
		Month: month,
	}, nil
}

func (r *statisticRepository) GetTodoStatistic(ctx context.Context, arg *repository.GetTodoStatisticParams) (*modeldb.TodoStatistic, error) {
	filter := bson.M{
		"userId": arg.UserID,
		"todoId": arg.TodoID,
	}

	result := r.collection.FindOne(ctx, filter)
	statistic := &modeldb.TodoStatistic{}
	err := result.Decode(statistic)
	if err != nil {
		return nil, mongodb.MongoErrorToDBError(err)
	}

	return statistic, nil
}

func (r *statisticRepository) CreateTodoStatistic(ctx context.Context, arg *repository.CreateTodoStatisticParams) (*modeldb.TodoStatistic, error) {
	statistic := &modeldb.TodoStatistic{
		TodoID: arg.TodoID,
		UserID: arg.UserID,
		Month:  arg.Month,
		Year:   arg.Year,
	}

	result, err := r.collection.InsertOne(ctx, statistic)
	if err != nil {
		return nil, mongodb.MongoErrorToDBError(err)
	}

	id, err := mongodb.InsertedIDToHex(result.InsertedID)
	if err != nil {
		return nil, err
	}

	return r.getTodoStatisticByID(ctx, id)
}

func (r *statisticRepository) getUserTodoStatisticMonthly(ctx context.Context, userID, month string) (int64, error) {
	filter := bson.M{
		"userId": userID,
		"month":  month,
	}

	return r.collection.CountDocuments(ctx, filter)
}

func (r *statisticRepository) getUserTodoStatisticYearly(ctx context.Context, userID, year string) (int64, error) {
	filter := bson.M{
		"userId": userID,
		"year":   year,
	}

	return r.collection.CountDocuments(ctx, filter)
}

func (r *statisticRepository) getStatisticTotal(ctx context.Context, userID string) (int64, error) {
	filter := bson.M{
		"userId": userID,
	}

	return r.collection.CountDocuments(ctx, filter)
}

func (r *statisticRepository) getTodoStatisticByID(ctx context.Context, id string) (*modeldb.TodoStatistic, error) {
	objectID, err := mongodb.IDHexToObjectID(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": objectID,
	}
	result := r.collection.FindOne(ctx, filter)
	statistic := &modeldb.TodoStatistic{}
	err = result.Decode(statistic)
	if err != nil {
		return nil, mongodb.MongoErrorToDBError(err)
	}

	return statistic, nil
}
