package main

import (
	"fmt"
	"math"
)

// Under the hood, interface values can be thought of as a tuple of a value and a concrete type:
// (value, type)
// An interface value holds a value of a specific underlying concrete type.
//
// Calling a method on an interface value executes the method of the same name on its underlying type.

type I1 interface {
	M1()
}

// - - - T1 - - -
type T1 struct {
	S string
}

func (t *T1) M1() {
	fmt.Println(t.S)
}

// - - - F1 - - -
type F1 float64

func (f F1) M1() {
	fmt.Println(f)
}

// - - -

func main() {
	var i I1

	i = &T1{"Hello"}
	describe(i)
	i.M1()

	i = F1(math.Pi)
	describe(i)
	i.M1()
}

func describe(i I1) {
	fmt.Printf("(%v, %T)\n", i, i)
}
