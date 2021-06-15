package main

import (
	"fmt"
	"math"
)

type Vertex3 struct {
	X, Y float64
}

func Abs3(v Vertex3) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Here we see the `Abs` and `Scale` methods rewritten as functions.
//
// Again, try removing the `*` from `Scale` method.
// Can you see why the behavior changes?
// What else did you need to change for the example to compile?
//
// (If you're not sure, continue to the next page.)
func Scale(v *Vertex3, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex3{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs3(v))
}
