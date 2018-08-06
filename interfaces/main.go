package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct {}
type spanishBot struct {}

func (englishBot) getGreeting() string{
	//custom logic for generating an english greeting

	return "hello there"
}
func (spanishBot) getGreeting() string{
	//custom logic for generating an spanish greeting

	return "hola!"
}

func printGreeting(b bot)  {
	fmt.Println(b.getGreeting())
}
func main() {
	eb:=englishBot{}
	sb:=spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}
