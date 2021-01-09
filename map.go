package main

import "fmt"

func mapExample() {
	// Map Declaration 1
	// var colors map[string]string

	// Declaration 2
	// colors := make(map[string]string)

	// Declaration 3
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4fg454",
		"brown": "#5fy410",
	}

	colors["white"] = "#ffffff"

	delete(colors, "green")

	printMap(colors)

	fmt.Println(colors)
}

func printMap(mapData map[string]string) {
	for key, value := range mapData {
		fmt.Println(key, value)
	}
}
