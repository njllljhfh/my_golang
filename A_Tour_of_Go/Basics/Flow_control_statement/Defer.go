package main

import "fmt"

// A defer statement defers the execution of a function until the surrounding function returns.
// The deferred call's arguments are evaluated immediately,
// but the function call is not executed until the surrounding function returns.
func fa() {
	defer fmt.Println("dragon")
	defer fmt.Println("is")
	fmt.Println("name")
}

func main() {
	fmt.Println("My")

	defer fmt.Println("world")

	fa()

	fmt.Println("hello")
}
