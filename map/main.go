package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "x00022",
		"green": "x00023",
		"white": "ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {

	delete(c, "white")
	for key, value := range c {
		fmt.Println("The Hex code for the color ", key, " is ", value)
	}
}
