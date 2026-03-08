package main

import "fmt"

func main() {
	// Declare and initialize an array
	var myArray = [5]int{10, 20, 30, 40, 50}

	// Iterate over the array and print each element
	for i, value := range myArray {
		fmt.Printf("Index: %d, Value: %d\n", i, value)
	}
}
