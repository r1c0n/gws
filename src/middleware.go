package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/r1c0n/gws/middleware"
)

func ApplyMiddleware(r *mux.Router, config Config) {
	middlewareHandlers := []func(http.Handler) http.Handler{}

	if config.Middleware.LoggingMiddlewareEnabled {
		middlewareHandlers = append(middlewareHandlers, middleware.LoggingMiddleware)
	}

	if config.Middleware.GzipMiddlewareEnabled {
		middlewareHandlers = append(middlewareHandlers, middleware.GzipMiddleware)
	}

	for _, mw := range middlewareHandlers {
		r.Use(mw)
	}
}
