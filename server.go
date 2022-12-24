package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	version = "1.0.0"
	author  = "Official B"
)

func main() {
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("html"))
	fmt.Print("Hello, World! The current version of gowebserver is " + version + " and it was written by " + author + ".\n-----------------------------\n")
	fmt.Print("To exit the program, enter the key combination \"CTRL + C\".\n")
	router.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	http.ListenAndServe(":8080", router)
}
