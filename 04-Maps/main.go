package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#008000",
		"white": "#FFFFFF",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for hex, color := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
