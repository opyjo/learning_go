package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi there!"
}

type spanishBot struct{}

func (spanishBot) getGreeting() string {
	return "Hola!"
}




type frenchBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	// fb := frenchBot{}

	printGreeting(eb)
	printGreeting(sb)
	// printGreeting(fb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }
