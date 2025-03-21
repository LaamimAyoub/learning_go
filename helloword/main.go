package hello

import (
	"fmt"
)

const helloword = "Hello "
const holla = "Hola "
const bonjour = "Bonjour "
const azul = "Azul "

const spanish = "Spanish"
const french = "French"
const Amazigh = "Amazigh"

func getGreeting(lang string) string {
	switch lang {
	case spanish:
		return holla
	case french:
		return bonjour
	case Amazigh:
		return azul

	default:
		return helloword
	}

}

func Hello(name string, lang string) string {

	greeting := getGreeting(lang)

	if name == "" {
		return greeting + "Nothing"
	}

	return greeting + name
}

func Main() {
	fmt.Println("Hello")
}
