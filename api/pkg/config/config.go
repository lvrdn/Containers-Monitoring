package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HTTPport   string `envconfig:"PORT"`
	DBhost     string `envconfig:"DB_HOST"`
	DBname     string `envconfig:"DB_NAME"`
	DBusername string `envconfig:"DB_USERNAME"`
	DBpassword string `envconfig:"DB_PASSWORD"`
}

func GetConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
