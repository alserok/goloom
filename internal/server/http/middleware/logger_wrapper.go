package middleware

import (
	"github.com/alserok/goloom/pkg/logger"
	"net/http"
)

func WithLogger(log logger.Logger) func(handlerFunc http.Handler) http.Handler {
	return func(handlerFunc http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(logger.WrapLogger(r.Context(), log))
			handlerFunc.ServeHTTP(w, r)
		})
	}
}
