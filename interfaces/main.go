package main

import "fmt"

type engslishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	// printGreeting(sb)

}
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (engslishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting

	return "Hi there!"
}

func (spanishBot) getGreeting() string {

	return "Hola!"
}
