package main

import "fmt"

// The empty interface
// The interface type that specifies zero methods is known as the `empty interface`:
// interface{}
//
// 空接口的重要功能：An empty interface may hold values of any type. (Every type implements at least zero methods.)
//
// Empty interfaces are used by code that handles values of unknown type.
// For example, `fmt.Print` takes any number of arguments of type interface{}.

func main() {
	var i interface{}
	describe4(i) // 注意：这里的输出(<nil>, <nil>), 空接口的value 和 type 都是 nil.

	i = 42
	describe4(i)

	i = "hello"
	describe4(i)
	fmt.Println("- - - - - - - - - - - - - - - - - - - - ")

	// dragon test
	dragon1(1)
	dragon1("a")
	array := []int{1, 2, 3}
	dragon1(array)
	dragon1(&array)
	type X1 struct {
		X, Y int
	}
	x1 := X1{3, 4}
	dragon1(x1)
	dragon1(&x1)
	// - - -
}

func dragon1(a ...interface{}) {
	// The argument is an array of empty interface.
	for _, item := range a {
		fmt.Printf("(%v, %T)\n", item, item)
	}
}

func describe4(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
