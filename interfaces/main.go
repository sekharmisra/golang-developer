package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}
type bot interface {
	getGreeting() string
}

func main() {

	eb := englishBot{}
	sp := spanishBot{}

	printGreeting(eb)
	printGreeting(sp)
}

func (eb englishBot) getGreeting() string {
	//Very custom logic for English Bot
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	//Very custom logic for Spanish Bot
	return "Hola!!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
