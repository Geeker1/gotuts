package main

import "fmt"
// Interfaces are good for code reuse

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
	fmt.Println(eb.getValue())
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}


// (e englishBot) === (englishBot) if we are not using the e value

func (englishBot) getValue() int {
	return 42
}

func (englishBot) getGreeting() string{
	// VERY custom logic for generating an english greeting
	return "Hi there"
}

func (spanishBot) getGreeting() string{
	return "Hola!"
}