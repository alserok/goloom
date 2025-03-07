package config

import (
	"os"
	"strings"
)

type Config struct {
	Env  string
	Port string

	Storage struct {
		Dir string
	}

	Services []string
}

func MustLoad() *Config {
	var cfg Config

	cfg.Env = os.Getenv("ENV")
	cfg.Port = os.Getenv("PORT")

	cfg.Storage.Dir = os.Getenv("DIR")

	if servicesEnv := os.Getenv("SERVICES"); servicesEnv != "" {
		cfg.Services = strings.Split(servicesEnv, ",")
	}

	return &cfg
}
