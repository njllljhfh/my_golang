package main

import "golang.org/x/tour/reader"

// Exercise: Readers
// Implement a `Reader` type that emits an infinite stream of the ASCII character 'A'.
type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (m MyReader) Read(b []byte) (n int, err error) {
	// 将该函数复制到官网教程中执行(https://tour.golang.org/methods/22)
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
