package server

import (
	"github.com/alserok/goloom/internal/server/http"
	"github.com/alserok/goloom/internal/service"
	"github.com/alserok/goloom/pkg/logger"
)

type Server interface {
	MustServe(port string)
	Shutdown() error
}

const (
	HTTP = iota
)

func New(t uint, srvc service.Service, log logger.Logger) Server {
	switch t {
	case HTTP:
		return http.NewServer(srvc, log)
	default:
		panic("invalid server type")
	}
}
