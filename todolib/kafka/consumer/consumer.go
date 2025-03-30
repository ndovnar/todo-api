package consumer

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/rs/zerolog/log"
)

type Consumer struct {
	consumer *kafka.Consumer
	options  *Options
	stopChan chan any
}

func NewConsumer(bootstrapServers, topics []string, groupID string, opts ...Option) (*Consumer, error) {
	options := &Options{
		bootstrapServers: bootstrapServers,
		topics:           topics,
		groupID:          groupID,
		sessionTimeoutMs: 30000,
	}

	for _, o := range opts {
		o(options)
	}

	if len(options.bootstrapServers) == 0 {
		return nil, errors.New("bootstrap servers is required")
	}

	if len(options.topics) == 0 {
		return nil, errors.New("topics is required")
	}

	if options.groupID == "" {
		return nil, errors.New("group ID is required")
	}

	c := &Consumer{
		options:  options,
		stopChan: make(chan any),
	}

	var err error
	c.consumer, err = kafka.NewConsumer(optionsToMap(options))
	if err != nil {
		return nil, err
	}

	return c, nil
}

func optionsToMap(options *Options) *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers":  strings.Join(options.bootstrapServers, ","),
		"group.id":           options.groupID,
		"session.timeout.ms": options.sessionTimeoutMs,
	}
}

func (c *Consumer) Start(handler func(*kafka.Message)) error {
	if err := c.consumer.SubscribeTopics(c.options.topics, c.options.rebalanceCb); err != nil {
		return fmt.Errorf("failed to subscribe to topics: %w", err)
	}

	for {
		select {
		case <-c.stopChan:
			log.Info().Msg("consumer received stop signal")
			return nil
		default:
			msg, err := c.consumer.ReadMessage(1 * time.Second)
			if err != nil {
				if err.(kafka.Error).Code() == kafka.ErrTimedOut {
					continue
				}

				log.Error().Msgf("error reading message: %v", err)
				continue
			}

			if msg != nil {
				handler(msg)
			}
		}
	}
}

func (c *Consumer) RunShutdown(ctx context.Context) error {
	<-ctx.Done()
	close(c.stopChan)
	// wait a bit for the consumer loop to exit
	time.Sleep(2 * time.Second)

	return c.consumer.Close()
}
