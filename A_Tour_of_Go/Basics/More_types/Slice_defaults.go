package main

import "fmt"

// When slicing, you may omit the high or low bounds to use their defaults instead.
// The default is zero for the low bound and the length of the slice for the high bound.
func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	// For the array `s` these slice expressions are equivalent:
	s = s[0:6]
	fmt.Println(s)

	s = s[:6]
	fmt.Println(s)

	s = s[0:]
	fmt.Println(s)

	s = s[:]
	fmt.Println(s)
}
