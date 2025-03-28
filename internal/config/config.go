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
		RootDir string
		Dirs    []string
	}

	ServicesCheckPeriod time.Duration
}

func MustLoad() *Config {
	var cfg Config

	cfg.Env = os.Getenv("ENV")
	cfg.Port = os.Getenv("PORT")

	cfg.Storage.RootDir = os.Getenv("ROOT_DIR")
	cfg.Storage.Dirs = strings.Split(os.Getenv("DIRS"), ";")

	cfg.ServicesCheckPeriod = defaultServicesCheckPeriod
	if checkperiod, err := time.ParseDuration(os.Getenv("SERVICES_CHECK_PERIOD")); err == nil {
		cfg.ServicesCheckPeriod = checkperiod
	}

	return &cfg
}

const (
	defaultServicesCheckPeriod = time.Minute
)
