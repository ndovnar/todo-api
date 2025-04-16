package service

import (
	"context"
	"errors"
	"fmt"
	"time"
	"todolib/db"

	"todostatistic/internal/converters"
	"todostatistic/internal/model"
	"todostatistic/internal/repository"
)

type StatisticService interface {
	GetUserTodoStatistic(ctx context.Context, userID string) (*model.UserTodoStatistic, error)
	CreateTodoStatistic(ctx context.Context, arg *CreateStatisticParams) (*model.TodoStatistic, error)
}

type statisticService struct {
	statisticRepository repository.StatisticRepository
}

func NewStatisticService(statisticRepository repository.StatisticRepository) StatisticService {
	return &statisticService{
		statisticRepository: statisticRepository,
	}
}

func (s *statisticService) GetUserTodoStatistic(ctx context.Context, userID string) (*model.UserTodoStatistic, error) {
	userTodoStatistic, err := s.statisticRepository.GetUserTodoStatistic(ctx, &repository.GetUserTodoStatisticParams{
		UserID: userID,
		Year:   getCurrentYear(),
		Month:  getCurrentMonth(),
	})
	if err != nil {
		return nil, err
	}

	return converters.UserTodoStatisticModelDBToModel(userTodoStatistic), nil
}

type CreateStatisticParams struct {
	UserID string
	TodoID string
}

func (s *statisticService) CreateTodoStatistic(ctx context.Context, arg *CreateStatisticParams) (*model.TodoStatistic, error) {
	currentYear := getCurrentYear()
	currentMonth := getCurrentMonth()

	statistic, err := s.statisticRepository.GetTodoStatistic(ctx, &repository.GetTodoStatisticParams{
		UserID: arg.UserID,
		TodoID: arg.TodoID,
	})
	if err != nil && !errors.Is(err, db.ErrNotFound) {
		return nil, err
	}

	if statistic != nil {
		return converters.TodoStatisticModelDBToModel(statistic), nil
	}

	newStatistic, err := s.statisticRepository.CreateTodoStatistic(ctx, &repository.CreateTodoStatisticParams{
		UserID: arg.UserID,
		TodoID: arg.TodoID,
		Year:   currentYear,
		Month:  currentMonth,
	})
	if err != nil {
		return nil, err
	}

	return converters.TodoStatisticModelDBToModel(newStatistic), nil
}

func getCurrentYear() string {
	return fmt.Sprintf("%d", time.Now().Year())
}

func getCurrentMonth() string {
	return time.Now().Format("01")
}
