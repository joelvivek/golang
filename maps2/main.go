package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#473463",
		"white": "3fff62",
	}

	// var colors map[string]string

	fmt.Println("colors in map", colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex code for ", color, "is ", hex)
	}
}
