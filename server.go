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
	Port      string `json:"port"`       // Port is the port number that the server will listen on.
	StaticDir string `json:"static_dir"` // StaticDir is the directory where static assets (e.g. HTML, CSS, JavaScript) are stored.

	// TLSConfig contains configuration options for TLS (Transport Layer Security).
	TLSConfig struct {
		CertFile string `json:"cert_file"` // CertFile is the path to the TLS certificate file.
		KeyFile  string `json:"key_file"`  // KeyFile is the path to the TLS key file.
	} `json:"tls_config"`

	// RepoConfig contains information about the repository.
	RepoConfig struct {
		Version    string `json:"version"`    // Version is the version of the repository.
		Author     string `json:"author"`     // Author is the author of the repository.
		Product    string `json:"product"`    // Product is the name of the product that the repository is for.
		Repository string `json:"repository"` // Repository is the name of the repository.
	} `json:"repo_config"`
}

func main() {
	// Read the configuration file
	configData, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err) // If there is an error reading the file, log the error and exit the program.
	}

	// Parse the configuration data into a Config struct.
	var config Config
	if err := json.Unmarshal(configData, &config); err != nil {
		log.Fatal(err) // If there is an error parsing the data, log the error and exit the program.
	}

	// Clear the screen.
	screen.Clear()

	// Print a message coontaining information from the Config struct. Mostly repository information.
	fmt.Print("Hello, World! | ", config.RepoConfig.Product, " v", config.RepoConfig.Version, " | Created by ", config.RepoConfig.Author)
	fmt.Print("\nTo contribute, check out our GitHub repo: ", config.RepoConfig.Repository, ".")
	fmt.Print("\n----------------------------------------------------------------------------\n")
	fmt.Print("To exit the program, enter the key combination \"CTRL + C\".\n")
	fmt.Print("Site URL: http://localhost", config.Port, "\n")

	r := mux.NewRouter() // Create a new router.

	// Set up a FileServer to serve static assets.
	fs := http.FileServer(http.Dir(config.StaticDir))
	r.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	// Start the server.
	http.ListenAndServe(config.Port, r)
}
