package main

import "fmt"

// Let's have some fun with functions.
// Implement a `fibonacci` function that returns a function (a closure)
// that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a := 0
	b := 0
	return func() int {
		if a == 0 && b == 0 {
			a = 1
			return 0
		} else if b == 0 {
			a = 0
			b = 1
			return 1
		} else {
			a, b = b, a+b
			return b
		}
	}
}

func Fibonacci() func() int {
	a, b := -1, 1
	return func() int {
		if a == -1 {
			a = 0
			return a
		}
		a, b = b, a+b
		return a
	}
}

func main() {
	f := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
