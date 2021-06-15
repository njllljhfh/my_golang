package main

import (
	"fmt"
	"math"
)

type MyFloat float64

// You can declare a method on non-struct types, too.
//
// In this example we see a numeric type MyFloat with an Abs method.
//
// 注意：You can only declare a method with a receiver
// whose `type` is defined in the same package as the method.
// You cannot declare a method with a receiver
// whose `type` is defined in another package (which includes the built-in types such as `int`).
func (f MyFloat) Abs() float64 {
	if f < 0 {
		// MyFloat 与 float64 不能直接视为同类型，需要进行类型转换，因为Abs()方法返回的是 float64类型。
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f)
	fmt.Println(f.Abs())
}
