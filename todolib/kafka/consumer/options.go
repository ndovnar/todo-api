package consumer

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

type Options struct {
	bootstrapServers []string
	topics           []string
	groupID          string
	sessionTimeoutMs int
	rebalanceCb      kafka.RebalanceCb
}

type Option func(*Options)

func OptionTimeoutMs(timeout int) Option {
	return func(o *Options) {
		o.sessionTimeoutMs = timeout
	}
}

func OptionRebalanceCb(rebalanceCb kafka.RebalanceCb) Option {
	return func(o *Options) {
		o.rebalanceCb = rebalanceCb
	}
}
