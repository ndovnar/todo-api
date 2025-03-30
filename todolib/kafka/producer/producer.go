package producer

import (
	"errors"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Producer struct {
	producer *kafka.Producer
	options  *Options
}

const flushTimeout = 100

var (
	ErrUnknownEventType = errors.New("unknown event type")
)

func NewProducer(bootstrapServers []string, clientID string, opts ...Option) (*Producer, error) {
	options := &Options{
		bootstrapServers: bootstrapServers,
		clientID:         clientID,
		messageTimeoutMs: 5000,
	}

	for _, o := range opts {
		o(options)
	}

	if len(options.bootstrapServers) == 0 {
		return nil, errors.New("bootstrap servers is required")
	}

	if options.clientID == "" {
		return nil, errors.New("client id is required")
	}

	p := &Producer{
		options: options,
	}

	var err error
	p.producer, err = kafka.NewProducer(optionsToMap(options))
	if err != nil {
		return nil, err
	}

	return p, nil
}

func optionsToMap(options *Options) *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers":  strings.Join(options.bootstrapServers, ","),
		"client.id":          options.clientID,
		"message.timeout.ms": options.messageTimeoutMs,
	}
}

func (p *Producer) Produce(topic string, key, value []byte) error {
	kafkaMsg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value:     value,
		Key:       key,
		Timestamp: time.Now(),
	}
	kafkaChan := make(chan kafka.Event)
	if err := p.producer.Produce(kafkaMsg, kafkaChan); err != nil {
		return err
	}

	e := <-kafkaChan
	switch ev := e.(type) {
	case *kafka.Message:
		return nil
	case kafka.Error:
		return ev
	default:
		return ErrUnknownEventType
	}
}

func (p *Producer) Close() {
	p.producer.Flush(flushTimeout)
	p.producer.Close()
}
