package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/r1c0n/gws/middleware"
)

func startServer(config Config) {
	r := mux.NewRouter()

	if config.Middleware.LoggingMiddlewareEnabled && config.Middleware.GzipMiddlewareEnabled {
		r.Use(middleware.LoggingMiddleware, middleware.GzipMiddleware)
	} else {
		switch {
		case config.Middleware.LoggingMiddlewareEnabled:
			r.Use(middleware.LoggingMiddleware)
		case config.Middleware.GzipMiddlewareEnabled:
			r.Use(middleware.GzipMiddleware)
		}
	}

	//r.Use(middleware.LoggingMiddleware, middleware.GzipMiddleware)

	fs := http.FileServer(http.Dir(config.StaticDir))
	r.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	if config.TLSConfig.Enabled {
		server := &http.Server{
			Addr:    config.Port,
			Handler: r,
			TLSConfig: &tls.Config{
				Certificates: getCertificates(config.TLSConfig.CertFile, config.TLSConfig.KeyFile),
			},
		}

		log.Fatal(server.ListenAndServeTLS("", ""))
	} else {
		log.Fatal(http.ListenAndServe(config.Port, r))
	}
}
