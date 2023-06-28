package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/inancgumus/screen"
)

type Config struct {
	Port      string `json:"port"`
	Domain    string `json:"domain"`
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
	config := loadConfig()

	printHeader(config)
	startServer(config)
}

func loadConfig() Config {
	configData, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	if err := json.Unmarshal(configData, &config); err != nil {
		log.Fatal(err)
	}

	return config
}

func printHeader(config Config) {
	screen.Clear()
	fmt.Printf("Hello, World! | %s v%s | Created by %s\n", config.RepoConfig.Product, config.RepoConfig.Version, config.RepoConfig.Author)
	fmt.Printf("To contribute, check out our GitHub repo: %s\n", config.RepoConfig.Repository)
	fmt.Println("----------------------------------------------------------------------------")
	fmt.Printf("To exit the program, enter the key combination \"CTRL + C\".\n")
	fmt.Printf("Site URL: http://%s%s\n", config.Domain, config.Port)
}
