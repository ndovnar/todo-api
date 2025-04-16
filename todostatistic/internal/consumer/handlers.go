package consumer

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"todostatistic/internal/consumer/handlers"
	"todostatistic/internal/service"
)

type Handlers struct {
	handlers []Handler
}

type Handler func(msg *kafka.Message)

func (h *Handlers) HandleMessage(msg *kafka.Message) {
	for _, handler := range h.handlers {
		handler(msg)
	}
}

func (h *Handlers) on(topic string, handler Handler) {
	topicHandler := func(msg *kafka.Message) {
		if topic == *msg.TopicPartition.Topic {
			handler(msg)
		}
	}

	h.handlers = append(h.handlers, topicHandler)
}

func (h *Handlers) RegisterHandlers(todoEventsTopic string, statisticService service.StatisticService) {
	todoHandlers := handlers.NewTodo(statisticService)

	h.on(todoEventsTopic, todoHandlers.HandleTodoEvent)
}
