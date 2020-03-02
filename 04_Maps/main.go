package main

import "fmt"

func main() {

	// initializing to it's zero value
	var colors map[string]string

	//another way to initialize a map
	//colors:= make( map[string]string )

	colors = map[string]string{
		"red":   "#ff0000",
		"green": "#00aa00 ",
	}

	// add new value
	colors["white"] = "#ffffff"

	// replace existing value
	colors["green"] = "#00ff00"

	// delete from map
	delete(colors, "white")

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
