package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z) // Newton's method.
		fmt.Printf("Sqrt(%d)---%g\n", i, z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
