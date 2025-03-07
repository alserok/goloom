package server

import (
	"github.com/alserok/goloom/internal/server/http"
	"github.com/alserok/goloom/internal/service"
)

type Server interface {
	MustServe(port string)
	Shutdown() error
}

const (
	HTTP = iota
)

func New(t uint, srvc service.Service) Server {
	switch t {
	case HTTP:
		return http.NewServer(srvc)
	default:
		panic("invalid server type")
	}
}
