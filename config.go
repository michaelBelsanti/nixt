package main

import (
	"bufio"
	"log"
	"os"

	"github.com/adrg/xdg"
)

func getFlakes() []string {
	configFile, err := xdg.SearchConfigFile("nixt/flakes")
	if err != nil {
		log.Println("Could not find configuration file, using defaults.")
		return []string{"github:nix-community/templates"}
	}

	file, err := os.Open(configFile)
	if err != nil {
		log.Println("Failed to open configuration file, using defaults.")
	}

	defer file.Close()

	var flakes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		flakes = append(flakes, scanner.Text())
	}
	return flakes
}
