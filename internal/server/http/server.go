package http

import (
	"context"
	"errors"
	"github.com/alserok/goloom/internal/service"
	"net"
	"net/http"
)

func NewServer(srvc service.Service) *server {
	mux := http.NewServeMux()

	setupRoutes(mux, newHandler(srvc))

	return &server{
		s: &http.Server{Handler: mux},
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
