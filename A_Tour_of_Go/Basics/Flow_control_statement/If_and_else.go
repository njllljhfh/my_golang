package main

import (
	"fmt"
	"math"
)

func pow1(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else if v < 30 {
		fmt.Printf("%g <= 30\n", v)
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

// Both calls to `pow1` return their results before the call to `fmt.Println` in `main` begins.
func main() {
	fmt.Println(
		pow1(3, 2, 10),
		pow1(3, 3, 20),
		pow1(3, 4, 20),
	)
}
