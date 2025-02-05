package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Addresses     []string
	PingTimeout   time.Duration
	PingFrequency time.Duration
	AddrAPI       string
	MethodAPI     string
}

func GetConfig(path string) (*Config, error) {
	file, err := os.Open("./app.env")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	env, err := godotenv.Parse(file)
	if err != nil {
		return nil, err
	}

	timeout, err := strconv.Atoi(env["TIMEOUT"])
	if err != nil {
		return nil, err
	}

	frequency, err := strconv.Atoi(env["FREQUENCY"])
	if err != nil {
		return nil, err
	}

	addresses := strings.Split(env["ADDRESSES"], ",")
	if len(addresses) == 0 {
		return nil, fmt.Errorf("no addresses to ping, check cfg file")
	}

	return &Config{
		Addresses:     addresses,
		PingTimeout:   time.Millisecond * time.Duration(timeout),
		PingFrequency: time.Millisecond * time.Duration(frequency),
		AddrAPI:       env["ADDRESS_API"],
		MethodAPI:     env["METHOD_API"],
	}, nil
}
