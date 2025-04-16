package consumer

type Config struct {
	BootstrapServers []string `required:"true"`
	TodoEventsTopic string `required:"true"`
	GroupID          string   `default:"todostatistic"`
}
