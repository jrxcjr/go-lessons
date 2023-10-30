package main

import "fmt"

func main() {

	//initializing an empty map
	//var colors map[string]string

	//This is another way of initializing an empty map
	//colors := make(map[string]string)

	//assignment of new key-value pair to an existing map
	//colors["white"] = "#fffff"

	//deleting a key from a value is done by using the delete function
	//delete(colors, "white")

	//basic map syntax
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#4bf745",
		"white": "#fffff",
	}

	printMap(colors)
}

// the argument name is followed by the type of the map that is being iterated
func printMap(c map[string]string) {
	//this receives two variables from the argument (c), color (key) and hex (value)
	for color, hex := range c {
		fmt.Println("hex code for", color, "is", hex)
	}

}
