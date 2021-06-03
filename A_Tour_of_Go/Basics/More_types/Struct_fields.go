package main

import "fmt"

type Vertex1 struct {
	X int
	Y int
}

func main() {
	v := Vertex1{1, 2}
	// Struct fields are accessed using a dot.
	v.X = 4
	fmt.Println(v.X)
}
