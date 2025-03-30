package consumer

import (
	"context"
	"fmt"

	"todolib/kafka/consumer"

	"todostatistic/internal/service"
)

type Consumer struct {
	consumer *consumer.Consumer
	handlers *Handlers
}

func NewConsumer(config *Config, statisticService service.StatisticService) (*Consumer, error) {
	topics := []string{config.TodoEventsTopic}
	consumer, err := consumer.NewConsumer(config.BootstrapServers, topics, config.GroupID)
	if err != nil {
		return nil, fmt.Errorf("failed to create kafka consumer, %w", err)
	}

	handlers := &Handlers{}
	handlers.RegisterHandlers(config.TodoEventsTopic, statisticService)

	return &Consumer{
		consumer: consumer,
		handlers: handlers,
	}, nil
}

func (c *Consumer) Run() error {
	return c.consumer.Start(c.handlers.HandleMessage)
}

func (c *Consumer) RunShutdown(ctx context.Context) error {
	return c.consumer.RunShutdown(ctx)
}
