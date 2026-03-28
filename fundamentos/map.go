package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	delete(colors, "green")

	colors["yellow"] = "#ffff00"

	if value, exists := colors["white"]; exists {
		fmt.Println("White color exists:", value)
	} else {
		fmt.Println("White color does not exist")
	}

	for key, valur := range colors {
		fmt.Printf("%s: %s\n", key, valur)
	}
}
