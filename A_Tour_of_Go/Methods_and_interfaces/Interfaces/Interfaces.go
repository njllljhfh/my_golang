package main

import (
	"fmt"
	"math"
)

// An `interface` type is defined as a set of method signatures.
type Abser interface {
	Abs1() float64
}

func main() {
	// A value of interface type can hold any value that implements those methods.
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// Note: There is an error in the example code(a = v).
	// `Vertex` (the value type) doesn't implement `Abser`
	// because the Abs method is defined only on `*Vertex` (the pointer type).
	//
	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	a = v

	fmt.Println(a.Abs1())
}

type MyFloat float64

func (f MyFloat) Abs1() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs1() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
