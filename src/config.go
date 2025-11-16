package main

import (
	"encoding/json"
	"log"
	"os"
)

var (
	version string = "1.5.0"
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
	CORS struct {
		Enabled          bool     `json:"enabled"`
		AllowedOrigins   []string `json:"allowed_origins"`
		AllowedMethods   []string `json:"allowed_methods"`
		AllowedHeaders   []string `json:"allowed_headers"`
		AllowCredentials bool     `json:"allow_credentials"`
		MaxAge           int      `json:"max_age"`
	} `json:"cors"`
	RateLimit struct {
		Enabled           bool     `json:"enabled"`
		RequestsPerMinute int      `json:"requests_per_minute"`
		Burst             int      `json:"burst"`
		Whitelist         []string `json:"whitelist"`
		ExemptPaths       []string `json:"exempt_paths"`
	} `json:"rate_limit"`
}

func loadConfig(filename string) Config {
	configData, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read config file '%s': %v", filename, err)
	}

	var config Config
	if err := json.Unmarshal(configData, &config); err != nil {
		log.Fatalf("Failed to parse config file '%s': %v", filename, err)
	}

	return config
}
