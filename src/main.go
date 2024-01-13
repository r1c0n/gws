package main

func main() {
	config := loadConfig("config.json")

	initializeUI(config)
	startServer(config)
}
