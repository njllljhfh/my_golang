package main

import (
	"fmt"
	"math"
)

// Unlike in C, in Go assignment between items of different type requires an explicit conversion.
// Try removing the float64 or uint conversions in the example and see what happens.
func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	// var x, y int = -3, -4
	// var f float64 = float64(x + y)
	var z uint = uint(f)

	fmt.Printf("%T %T %T %T\n", x, y, f, z)
	fmt.Println(x, y, f, z)
}
