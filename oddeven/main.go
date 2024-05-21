package main

import "fmt"

func main() {

	//Declare slice of int from 0-10
	i := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	for _, value := range i {
		if value%2 == 0 {
			fmt.Printf("Number %v is Even\n", value)
		} else {
			fmt.Printf("Number %v is Odd\n", value)
		}
	}
}
