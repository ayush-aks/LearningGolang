package main

import "fmt"

func main() {
	// First method:
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#ff00ff",
		"white": "#fffff",
		"black": "#0a000",
	}
	// Second method:
	// colors := make(map[int]string) OR 	var colors map[string]string
	// colors[10] = "#fffff"

	delete(colors, "black")
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, "is ", hex)
	}
}
