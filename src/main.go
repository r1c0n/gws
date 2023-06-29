package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
	"github.com/inancgumus/screen"
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

func main() {
	config := loadConfig("config.json")

	printHeader(config)
	openURL(config)
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

func openURL(config Config) {
	var url string

	if config.TLSConfig.Enabled {
		url = fmt.Sprintf("https://%s%s", config.Domain, config.Port)
	} else {
		url = fmt.Sprintf("http://%s%s", config.Domain, config.Port)
	}

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		log.Fatalf("Unsupported platform")
	}

	if err := cmd.Start(); err != nil {
		log.Fatalf("Failed to open URL: %v", err)
	}
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
