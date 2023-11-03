package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func startServer(config Config) {
	r := mux.NewRouter()

	ApplyMiddleware(r, config)

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
