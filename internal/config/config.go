package config

import (
	"os"
	"strings"
	"time"
)

type Config struct {
	Env  string
	Port string

	Storage struct {
		Dir string
	}
	State struct {
		CheckPeriod time.Duration
		Services    []string
	}
}

func MustLoad() *Config {
	var cfg Config

	cfg.Env = os.Getenv("ENV")
	cfg.Port = os.Getenv("PORT")

	cfg.Storage.Dir = os.Getenv("DIR")

	if servicesEnv := os.Getenv("SERVICES"); servicesEnv != "" {
		cfg.State.Services = strings.Split(servicesEnv, ",")

		checkPeriod, err := time.ParseDuration(os.Getenv("CHECK_PERIOD"))
		if err == nil {
			cfg.State.CheckPeriod = checkPeriod
		} else {

		}
	}

	return &cfg
}
