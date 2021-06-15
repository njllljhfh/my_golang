package main

import "fmt"

type I2 interface {
	M2()
}

type T2 struct {
	S string
}

// Interface values with nil underlying values

// If the concrete value inside the interface itself is nil,
// the method will be called with a nil receiver.
//
// In some languages this would trigger a null pointer exception,
// but in Go it is common to write methods
// that gracefully handle being called with a nil receiver
// (as with the method `M` in this example.)
//
// Note that an interface value that holds a nil concrete value is itself non-nil. (这句话不理解)

func (t *T2) M2() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I2

	var t *T2
	i = t
	describe2(i) // 注意：这行输出的结果，interface 的 value 是 `nil`，但是 type 不是 `nil` 而是 `*main.T2`.
	i.M2()
	fmt.Println("- - - - - - - - - - - - - - - - - - - - ")

	i = &T2{"hello"}
	describe2(i)
	i.M2()
}

func describe2(i I2) {
	fmt.Printf("(%v, %T)\n", i, i)
}
