package main

func main() {
	config := loadConfig("config.json")

	printHeader(config)
	openURL(config)
	startServer(config)
}
