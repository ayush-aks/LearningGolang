package main

import "fmt"

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
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	// We are not making use of variable for struct englishBot, therefore we ommited the variable in receiver.
	return "Hello"
}

func (spanishBot) getGreeting() string {
	//VERY custom logic for generating a spanish greeting
	return "Hola!"
	// see interface type... Now spanishBot is pretend member of type bot(interface)
}
