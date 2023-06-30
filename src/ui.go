// ui.go
package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
	"github.com/inancgumus/screen"
)

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
