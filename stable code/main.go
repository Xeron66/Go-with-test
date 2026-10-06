package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return "Hola, " + name
	}
	return englishHelloPrefix + name
}

func main() {
	/*
		For both Hello("") and Hello('Chris') are called the test should pass.
	*/
	fmt.Println(Hello("Chris", "English"))
}
