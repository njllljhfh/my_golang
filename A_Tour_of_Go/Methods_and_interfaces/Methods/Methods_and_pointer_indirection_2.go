package main

import (
	"fmt"
	"math"
)

type Vertex5 struct {
	X, Y float64
}

func (v Vertex5) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex5) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex5{3, 4}
	fmt.Println(v.Abs())
	// The equivalent thing happens in the reverse direction.
	// Functions that take a value argument must take a value of that specific type:
	// var v Vertex
	// fmt.Println(AbsFunc(v))  // OK
	// fmt.Println(AbsFunc(&v)) // Compile error!
	fmt.Println(AbsFunc(v))
	fmt.Println("- - - - - - - - - - - - - - - - - - - - ")

	p := &Vertex5{4, 3}
	// while methods with value receivers take either a value or a pointer as the receiver when they are called:
	// var v Vertex
	// fmt.Println(v.Abs()) // OK
	// p := &v
	// fmt.Println(p.Abs()) // OK
	// 注意：In this case, the method call `p.Abs()` is interpreted as `(*p).Abs()`.
	fmt.Println(p.Abs())
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(*p))
}
