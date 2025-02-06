package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Addresses     []string      `envconfig:"ADDRESSES"`
	PingTimeout   time.Duration `envconfig:"TIMEOUT"`
	PingFrequency time.Duration `envconfig:"FREQUENCY"`
	AddrAPI       string        `envconfig:"ADDRESS_API"`
	MethodAPI     string        `envconfig:"METHOD_API"`
}

func GetConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
