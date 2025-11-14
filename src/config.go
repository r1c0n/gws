package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var (
	version string = "1.4.1"
	author  string = "recon (recon@mail.recon.best)"
	title   string = "Gamma Web Server"
	repo    string = "https://github.com/r1c0n/gws"
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
	Middleware struct {
		LoggingMiddlewareEnabled bool `json:"logging_middleware_enabled"`
		GzipMiddlewareEnabled    bool `json:"gzip_middleware_enabled"`
	} `json:"middleware"`
	ErrorPages struct {
		Enabled       bool              `json:"enabled"`
		ErrorPagesDir string            `json:"error_pages_dir"`
		Pages         map[string]string `json:"pages"`
	} `json:"error_pages"`
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
