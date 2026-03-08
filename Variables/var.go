package main

import "fmt"

// CREATING VARIABLES
// 1.
// using the var keyword to create a variable
// syntax: var <variable name> <variable type> = <value>
var card string = "Ace of Spades"
var number int = 10


//2.
// Grouping Variable Declarations
//syntax
// var (
//     variable1 variableType1
//     variable2 variableType2 = initialValue2
//     variable3 variableType3
// )

var (
	name     string
	age      int = 30
	location string
)

// 3.
// Short Declaration
// syntax: <variable name> := <value>
// only used when declaring a new variable
// cannot be used to reassign a variable
// only works inside a function
// := is not an operator, it is a statement
// := is shorthand for var <variable name> <variable type> = <value>
// cannot be used for package level variables
func main() {
	firstName := "Alex"
	fmt.Print(firstName)
	useComposite()
	useMake()	

}

// 4.
// Using the new keyword
// syntax: <variable name> := new(<variable type>)
// new keyword is used to create a pointer to a variable
// new keyword is used to create a variable of a type that is not primitive
var ptr *int = new(int)

// 5.
// Using composite literals
// syntax: <variable name> := <value>
// composite literals are used to create a value of a type that is not primitive
func useComposite() {
	numbers := []int{1, 2, 3} // slice literal
	person := struct {
		name string
		age  int
	}{"Alice", 30} // struct literal
	fmt.Print(numbers, person)
}

//6.
// Using 'make' keyword
// syntax: <variable name> := make(<type>, <size>)
// make keyword is used to create a slice, map, or channel only

func useMake() {
	numbers := make([]int, 0, 5)
	numbers = append(numbers, 1, 2, 3)
	fmt.Println(numbers)
}


