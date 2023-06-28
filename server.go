package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func startServer(config Config) {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir(config.StaticDir))
	r.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	http.ListenAndServe(config.Port, r)
}
