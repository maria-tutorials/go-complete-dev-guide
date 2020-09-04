package main

import "fmt"

func main() {
	// var colors map[string]string

	colors := make(map[string]string)

	/*
		colors := map[string]string{
			"red":   "#ff0000",
			"green": "#00ff00",
			"blue":  "#0000ff",
		}
	*/

	colors["white"] = "#ffffff"

	delete(colors, "white")

	fmt.Println(colors)
}
