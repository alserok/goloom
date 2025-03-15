package main

import (
	"github.com/alserok/goloom/internal/app"
	"github.com/alserok/goloom/internal/config"
)

func main() {
	app.MustStart(config.MustLoad())
}
