package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	/*
		For both Hello("") and Hello('Chris') are called the test should pass.
	*/
	fmt.Println(Hello("Chris"))
}
