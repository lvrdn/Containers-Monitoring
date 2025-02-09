package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HTTPport   string   `envconfig:"API_PORT"`
	DBhost     string   `envconfig:"DB_HOST"`
	DBname     string   `envconfig:"POSTGRES_DB"`
	DBusername string   `envconfig:"POSTGRES_USER"`
	DBpassword string   `envconfig:"POSTGRES_PASSWORD"`
	Addresses  []string `envconfig:"PING_ADDRESSES"`
}

func GetConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
