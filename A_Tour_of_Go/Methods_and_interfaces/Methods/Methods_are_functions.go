package main

import (
	"fmt"
	"math"
)

type Vertex1 struct {
	X, Y float64
}

// Remember: a method is just a function with a receiver argument.
// Here's Abs written as a regular function with no change in functionality.
func Abs(v Vertex1) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex1{3, 4}
	fmt.Println(Abs(v))
}
