package main

import (
	"encoding/json"
	"log"
	"os/exec"
)

type Template struct {
	Name        string
	Description string `json:"description"`
	Source      string
}

type ShowOutput struct {
	Templates map[string]Template `json:"templates"`
}

func getTemplates() []Template {
	var templates []Template
	flakes := getFlakes()

	for _, flake := range flakes {
		flakeTemplates := getTemplatesFromFlake(flake)
		templates = append(templates, flakeTemplates...)
	}

	return templates
}

func getTemplatesFromFlake(flake string) []Template {
	var templates []Template

	cmd := exec.Command("nix", "flake", "show", flake, "--json")
	output, err := cmd.Output()
	// jsonOutput := string(output)
	if err != nil {
		log.Fatal("Nix command failed: ", err)
	}

	var parsedOutput ShowOutput

	err = json.Unmarshal(output, &parsedOutput)
	if err != nil {
		log.Fatal("Failed to parse command output:\noutput")
	}

	for name, template := range parsedOutput.Templates {
		template.Name = name
		template.Source = flake
		templates = append(templates, template)
	}

	return templates
}
