package handlers

import (
	"context"
	"encoding/json"
	"todolib/kafka/events"
	"todostatistic/internal/service"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/rs/zerolog/log"
)

type Todo struct {
	statisticService service.StatisticService
}

func NewTodo(statisticService service.StatisticService) *Todo {
	return &Todo{
		statisticService: statisticService,
	}
}

func (h *Todo) HandleTodoEvent(msg *kafka.Message) {
	var todoEvent events.TodoEvent
	if err := json.Unmarshal(msg.Value, &todoEvent); err != nil {
		log.Error().Err(err).Msg("failed to unmarshal todo event")
		return
	}

	switch todoEvent.EventType {
	case events.TodoCreatedEventType:
		h.handleTodoCreatedEvent(&todoEvent)
	}
}

func (h *Todo) handleTodoCreatedEvent(todoEvent *events.TodoEvent) {
	ctx := context.Background()

	_, err := h.statisticService.CreateTodoStatistic(ctx, &service.CreateStatisticParams{
		UserID: todoEvent.UserID,
		TodoID: todoEvent.TodoID,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to create statistic")
	}
}
