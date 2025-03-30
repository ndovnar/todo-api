package producer

type Options struct {
	bootstrapServers []string
	clientID         string
	messageTimeoutMs int
}

type Option func(*Options)

func OptionMessageTimeoutMs(timeout int) Option {
	return func(o *Options) {
		o.messageTimeoutMs = timeout
	}
}
