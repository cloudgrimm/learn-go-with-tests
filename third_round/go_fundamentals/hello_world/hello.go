package main

import "fmt"

const englishHelloPrefix = "Hello, "

//Hello function that returns english greeting
func Hello(name, lang string) string {
	if lang == "Spanish" {
		return "Hola, " + name
	}
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Tinashe", "Spanish"))
}
