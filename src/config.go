package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Port      string `json:"port"`
	Domain    string `json:"domain"`
	StaticDir string `json:"static_dir"`
	TLSConfig struct {
		Enabled  bool   `json:"enabled"`
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

func loadConfig(filename string) Config {
	configData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	if err := json.Unmarshal(configData, &config); err != nil {
		log.Fatal(err)
	}

	return config
}
