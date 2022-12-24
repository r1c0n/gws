package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inancgumus/screen"
)

type Config struct {
	Port      string `json:"port"`
	StaticDir string `json:"static_dir"`

	TLSConfig struct {
		CertFile string `json:"cert_file"`
		KeyFile  string `json:"key_file"`
	} `json:"tls_config"`

	RepoConfig struct {
		Version    string `json:"version"`
		Author     string `json:"author"`
		Product    string `json:"product"`
		Repository string `json:"repository"`
	} `json:"repo_config"`
}

func main() {
	screen.Clear()
	router := mux.NewRouter()

	configData, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	if err := json.Unmarshal(configData, &config); err != nil {
		log.Fatal(err)
	}

	fmt.Print("Hello, World! | ", config.RepoConfig.Product, " v", config.RepoConfig.Version, " | Created by ", config.RepoConfig.Author)
	fmt.Print("\nTo contribute, check out our GitHub repo: ", config.RepoConfig.Repository, ".")
	fmt.Print("\n----------------------------------------------------------------------------\n")
	fmt.Print("To exit the program, enter the key combination \"CTRL + C\".\n")
	fmt.Print("Site URL: http://localhost", config.Port, "\n")

	fs := http.FileServer(http.Dir(config.StaticDir))
	router.PathPrefix("/").Handler(http.StripPrefix("/", fs))
	http.ListenAndServe(config.Port, router)
}
