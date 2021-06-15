package main

import "fmt"

type Vertex4 struct {
	X, Y float64
}

// while methods with pointer receivers take either a value or a pointer as the receiver when they are called:
// var v Vertex4
// v.Scale(5)  // OK
// p := &v
// p.Scale(10) // OK
//
// For the statement `v.Scale(5)`, even though `v` is a value and not a pointer,
// the method with the pointer receiver is called automatically.
// 注意：That is, as a convenience, Go interprets the statement `v.Scale(5)` as `(&v).Scale(5)`
// since the `Scale` method has a pointer receiver.
func (v *Vertex4) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// functions with a pointer argument must take a pointer:
// var v Vertex4
// ScaleFunc(v, 5)  // Compile error!
// ScaleFunc(&v, 5) // OK
func ScaleFunc(v *Vertex4, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex4{3, 4}
	v.Scale(2) // v is value receiver.
	ScaleFunc(&v, 10)

	p := &Vertex4{4, 3}
	p.Scale(3) // p is pointer receiver.
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
