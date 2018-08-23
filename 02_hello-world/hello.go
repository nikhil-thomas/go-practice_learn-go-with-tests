package main

import (
	"fmt"
	"math/rand"
)

const helloPrefix = "Hello, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const french = "French"
const frenchHelloPrefix = "Bonjour, "

// Hello returns "Hello, <p>"
func Hello(v, lang string) string {
	if v == "" {
		v = "World"
	}
	return greetingPrefix(lang) + v

}

func greetingPrefix(lang string) string {
	switch lang {
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenchHelloPrefix
	default:
		return helloPrefix

	}
}

func main() {
	fmt.Println(Hello("world", ""))

	rand.Intn(8)
}
