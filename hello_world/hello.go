package main

import (
	"fmt"
	"strings"
)

var language_hello  = map[string]string{
	"english": "Hello",
	"spanish": "Hola",
	"french": "Bonjour",
}

func Hello(name string, language string) string {
	language = strings.ToLower(language)
	greeting, ok := language_hello[language]
	if !ok {
		greeting = "Hello"
	}
	if name == "" {
		return greeting + ", World"
	}
	return greeting + ", " + name
}

func main() {
	fmt.Println(Hello("Anna", ""))
}
