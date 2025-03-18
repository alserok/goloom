package config

import (
	"os"
	"time"
)

type Config struct {
	Env  string
	Port string

	Storage struct {
		Dir string
	}

	ServicesCheckPeriod time.Duration
}

func MustLoad() *Config {
	var cfg Config

	cfg.Env = os.Getenv("ENV")
	cfg.Port = os.Getenv("PORT")

	cfg.Storage.Dir = os.Getenv("DIR")

	cfg.ServicesCheckPeriod = defaultServicesCheckPeriod
	if checkperiod, err := time.ParseDuration(os.Getenv("SERVICES_CHECK_PERIOD")); err == nil {
		cfg.ServicesCheckPeriod = checkperiod
	}

	return &cfg
}

const (
	defaultServicesCheckPeriod = time.Minute
)
