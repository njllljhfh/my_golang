package main

import "fmt"

// Nil interface values
// A nil interface value holds neither value nor concrete type.
//
// Calling a method on a nil interface is a run-time error
// because there is no type inside the interface tuple to indicate which `concrete` method to call.
type I3 interface {
	M()
}

func main() {
	var i I3
	describe3(i)
	i.M()
}

func describe3(i I3) {
	fmt.Printf("(%v, %T)\n", i, i)
}
