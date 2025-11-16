package main

import "github.com/r1c0n/gws/middleware"

func main() {
	config := loadConfig("config.json")

	initializeUI(config)
	startServer(config)
	middleware.CloseLogFiles()
}
