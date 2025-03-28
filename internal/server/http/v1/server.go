package v1

import (
	"context"
	"errors"
	"github.com/alserok/goloom/internal/service"
	"github.com/alserok/goloom/pkg/logger"
	"net"
	"net/http"
)

func NewServer(srvc service.Service, log logger.Logger) *server {
	s := &http.Server{}

	mux := http.NewServeMux()
	s.Handler = mux

	setupRoutes(mux, s, newHandler(srvc), log)

	return &server{
		s: s,
	}
}

type server struct {
	s *http.Server
}

func (s *server) Shutdown() error {
	return s.s.Shutdown(context.Background())
}

func (s *server) MustServe(port string) {
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic("failed to listen: " + err.Error())
	}

	if err = s.s.Serve(l); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic("failed to serve: " + err.Error())
	}
}
