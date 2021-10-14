package main

import "fmt"

type bot interface {
	getGreeting() string
}

/* @implements bot */
type englishBot struct{}
/* @implements bot */
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreetingEN(eb)
	printGreetingSP(sb)

	printGreeting(eb)
	printGreeting(sb)

}

func (englishBot) getGreeting() string {
	return "Hello, there!"
}

func (spanishBot) getGreeting() string {
	return "Hola !!! La Cucaracha !!!"
}

func printGreetingEN(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreetingSP(es spanishBot) {
	fmt.Println(es.getGreeting())
}


func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
