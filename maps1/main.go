package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#473463",
	// }

	// var colors map[string]string

	colors := make(map[string]string)
	colors["white"] = "#t44794"

	colors1 := make(map[int]string)
	colors1[17] = "4254"

	delete(colors1, 17)

	fmt.Println("colors in map", colors, colors1)

}
