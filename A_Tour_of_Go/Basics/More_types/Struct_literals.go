package main

import "fmt"

type Vertex struct {
	X, Y int
}

// A struct literal denotes a newly allocated struct value by listing the values of its fields.
// You can list just a subset of fields by using the `Name:` syntax. (And the order of named fields is irrelevant.)
// The special prefix `&` returns a pointer to the struct value.
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{Y: 1}  // X:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, v2, v3, p)
	fmt.Printf("%T", p)
}
