package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Addresses     []string      `envconfig:"PING_ADDRESSES"`
	PingTimeout   time.Duration `envconfig:"PING_TIMEOUT"`
	PingFrequency time.Duration `envconfig:"PING_FREQUENCY"`
	AddrAPI       string        `envconfig:"ADDRESS_API"`
	MethodAPI     string        `envconfig:"METHOD_API"`
}

func GetConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)

	if err != nil {
		return nil, err
	}

	if cfg.PingFrequency <= cfg.PingTimeout {
		return nil, fmt.Errorf("frequency value must be more than timeout value")
	}

	return cfg, nil
}
