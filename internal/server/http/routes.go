package http

import (
	"github.com/alserok/goloom/internal/server/http/middleware"
	"net/http"
)

func setupRoutes(s *http.ServeMux, h *handler) {
	s.HandleFunc("PUT /config/:path", middleware.WithErrorHandler(h.UpdateConfig))
	s.HandleFunc("DELETE /config/:path", middleware.WithErrorHandler(h.DeleteConfig))
	s.HandleFunc("GET /config/:path", middleware.WithErrorHandler(h.GetConfig))

	s.HandleFunc("GET /web/config/:path", middleware.WithErrorHandler(h.GetConfigPage))
	s.HandleFunc("GET /web/state", middleware.WithErrorHandler(h.GetStatePage))
}
