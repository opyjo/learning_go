package main

import "fmt"

func main() {
	var a int = 42
	var p *int = &a // p is a pointer to an integer and holds the address of a

	fmt.Println("Value of a:", a)      // Output the value of a
	fmt.Println("Address of a:", &a)   // Output the address of a
	fmt.Println("Value of p:", p)      // Output the value of p, which is the address of a
	fmt.Println("Dereferenced p:", *p) // Output the value at the address p points to
}
