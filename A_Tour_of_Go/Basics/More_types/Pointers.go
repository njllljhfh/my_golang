package main

import "fmt"

// Go has pointers. A pointer holds the memory address of a value.

// The type `*T` is a pointer to a `T` value. Its zero value is `nil`.
// The `&` operator generates a pointer to its operand.
// The `*` operator denotes the pointer's underlying value.

// This is known as "dereferencing" or "indirecting".
// Unlike C, Go has no pointer arithmetic.
func main() {
	i, j := 42, 2701

	p := &i // point to i
	fmt.Printf("The type of p is %T\n", p)
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	// dragon
	x := &i
	y := &i
	fmt.Println(x == y)
	fmt.Println(*x)
	fmt.Println(*y)
	// dragon-end
}
