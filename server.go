package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("html"))
	router.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	http.ListenAndServe(":8080", router)
}
