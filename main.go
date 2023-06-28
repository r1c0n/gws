package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fatih/color"
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
	config := loadConfig("config.json")

	printHeader(config)
	startServer(config)
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

func printHeader(config Config) {
	screen.Clear()
	fmt.Printf("%sHello, World! | %s v%s | Created by %s%s\n",
		color.YellowString(""), color.GreenString(config.RepoConfig.Product),
		color.CyanString(config.RepoConfig.Version), color.MagentaString(config.RepoConfig.Author),
		color.YellowString(""))
	fmt.Printf("To contribute, check out our GitHub repo: %s\n",
		color.BlueString(config.RepoConfig.Repository))
	fmt.Println("----------------------------------------------------------------------------")
	fmt.Printf("To exit the program, enter the key combination %s\"CTRL + C\"%s.\n",
		color.RedString(""), color.YellowString(""))
	fmt.Printf("Site URL: %shttp://%s%s\n", color.WhiteString(""), config.Domain, config.Port)
}
