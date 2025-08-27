package main

/* This is a simple Go program that prints "Hello, World!" to the console. */

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Welcome to Go programming!")

	// Using function from an external package
	fmt.Println(quote.Hello())
	fmt.Println("----------------------------------------------------------------")
	fmt.Println(quote.Go())
	fmt.Println("----------------------------------------------------------------")
	fmt.Println(quote.Glass())
	fmt.Println("----------------------------------------------------------------")
	fmt.Println(quote.Opt())
}
