package main

import (
	"fmt"
	"io"
	"strings"
)

// Readers
// The io package specifies the io.Reader interface, which represents the read end of a stream of data.
//
// The Go standard library contains many implementations of this interface,
// including files, network connections, compressors, ciphers, and others.
//
// The `io.Reader` interface has a `Read` method:
// func (T) Read(b []byte) (n int, err error)

// `Read` populates the given byte slice with data
// and returns the number of bytes populated and an error value.
// It returns an `io.EOF` error when the stream ends.
//
// The example code creates a `strings.Reader` and consumes its output 8 bytes at a time.
func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v, err = %v, b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		fmt.Println("- - - - - - - - - - - - - - - - - - - - ")
		if err == io.EOF {
			break
		}
	}

	fmt.Println("\nDragon test- - - - - - - - - - - - - - - - - - - - ")
	// dragon: 测试命名返回值
	fmt.Println(x())
	// dragon: 测试函数的出入参数(a)、函数体重的参数(array)、返回的参数(array)是否是同一个内存地址。
	// 结论：都是不同的地址。
	var a []byte
	// a := make([]byte, 2)
	fmt.Printf("Addr a = %p\n", &a)
	res := y(a)
	fmt.Printf("Addr res = %p\n", &res)
	fmt.Println(&res == &a)
	fmt.Println(res)
	fmt.Printf("%s", res)
}

// dragon
func x() (m int, n *int) {
	m = 1
	// 这里 n 没有赋值，返回时，返回 n 对应类型的默认`空值`。
	return
}

func y(array []byte) []byte {
	fmt.Printf("Addr array = %p\n", &array)
	for i := range make([]int, 2) {
		array = append(array, 0x41+byte(i))
	}
	return array
}
