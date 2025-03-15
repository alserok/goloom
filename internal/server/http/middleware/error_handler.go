package middleware

import (
	"github.com/alserok/goloom/internal/utils"
	"github.com/alserok/goloom/pkg/logger"
	"net/http"
)

func WithErrorHandler(handler func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			msg, code := utils.ParseErrorToHTTP(err)

			if code == http.StatusInternalServerError {
				logger.UnwrapLogger(r.Context()).Error(err.Error())
			}

			w.WriteHeader(code)
			_, _ = w.Write([]byte(msg))
		}
	}
}
