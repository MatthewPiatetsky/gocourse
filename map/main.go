package main

import "fmt"

func printMap (c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

func main() {
	// colors := make(map[string]string)

	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf645",
		"white": "#ffffff",
	}
	// colors["white"] = "#ffffff"
	// delete(colors, "white")

	printMap(colors)
}