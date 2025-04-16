package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"todolib/mongodb"

	"todostatistic/internal/consumer"
	"todostatistic/internal/httpapi"
)

type Config struct {
	LogLevel  zerolog.Level `default:"info" desc:"Level for generated logs"`
	Mongo     mongodb.Config
	HTTPAPI   httpapi.Config
	PublicKey string `required:"true"`
	Consumer  consumer.Config
}

func Load() (Config, error) {
	cfg := Config{}
	err := envconfig.Process("CONFIG", &cfg)

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(cfg.LogLevel)

	return cfg, err
}
