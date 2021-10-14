package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}
	fmt.Println(colors)

	// var colors2 map[string]string
	// colors2["white"] = "#FFFFFF"
	// fmt.Println(colors2)

	colors3 := make(map[string]string)
	colors3["black"] = "#000000"
	colors3["white"] = "#FFFFFF"
	fmt.Println(colors3)

	delete(colors3, "white")
	fmt.Println(colors3)

	printMap(colors)

}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println(color, hex)
	}
}
