package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Loader struct {
	prefix string
}

func NewLoader() *Loader {
	const (
		defaultPrefix = "CONFIG"
	)

	cl := &Loader{
		prefix: defaultPrefix,
	}

	return cl
}

func (cl *Loader) Load(cfg any) error {
	if err := envconfig.Process(cl.prefix, cfg); err != nil {
		return fmt.Errorf("could not load config using prefix %s: %w", cl.prefix, err)
	}

	return nil
}
