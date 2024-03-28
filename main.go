package main

import ()

func main() {
	for _, template := range getTemplates() {
		println("Description:", template.Source)
		println("Name:", template.Name)
		println("Description:", template.Description)
		println()
	}
}
