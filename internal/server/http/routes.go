package http

import (
	"github.com/alserok/goloom/internal/server/http/middleware"
	"github.com/alserok/goloom/pkg/logger"
	"net/http"
)

func setupRoutes(mux *http.ServeMux, s *http.Server, h *handler, log logger.Logger) {
	s.Handler = middleware.WithLogger(log)(s.Handler)

	mux.HandleFunc("PUT /config/update/", middleware.WithErrorHandler(h.UpdateConfig))
	mux.HandleFunc("DELETE /config/delete/", middleware.WithErrorHandler(h.DeleteConfig))
	mux.HandleFunc("GET /config/get/", middleware.WithErrorHandler(h.GetConfig))

	mux.HandleFunc("GET /web/config/file/", middleware.WithErrorHandler(h.GetConfigPage))
	mux.HandleFunc("GET /web/config/dir/", middleware.WithErrorHandler(h.GetDirPage))
	mux.HandleFunc("GET /web/state", middleware.WithErrorHandler(h.GetStatePage))
}
