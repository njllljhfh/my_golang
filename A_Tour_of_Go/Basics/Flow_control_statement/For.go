package main

import "fmt"

func main() {
	sum := 0
	// The init statement will often be a short variable declaration,
	// and the variables declared there are visible only in the scope of the for statement.
	for i := 0; i < 10; i++ {
		sum += i
	}
	// fmt.Print(i) // undefined: i
	fmt.Println(sum)
}
