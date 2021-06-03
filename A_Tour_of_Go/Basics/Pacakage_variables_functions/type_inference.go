package main

import "fmt"

func main() {
	// But when the right hand side contains an untyped numeric constant,
	// the new variable may be an int, float64, or complex128
	// depending on the precision of the constant:
	v := 42        // change me!
	v1 := "dragon" // change me!
	v2 := 18.0     // change me!
	g := 0.867 + 0.5i
	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("v1 is of type %T\n", v1)
	fmt.Printf("v2 is of type %T\n", v2)
	fmt.Printf("g is of type %T\n", g)

	// When the right hand side of the declaration is typed, the new variable is of that same type:
	var i = true
	j := i
	fmt.Printf("i is of type %T\n", j)
}
