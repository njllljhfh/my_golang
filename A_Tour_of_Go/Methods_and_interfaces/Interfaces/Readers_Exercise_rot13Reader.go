package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// 官网(https://tour.golang.org/methods/23)

// Exercise: rot13Reader
// A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.
//
// For example, the gzip.NewReader function
// takes an io.Reader (a stream of compressed data)
// and returns a `*gzip.Reader` that also implements io.Reader (a stream of the decompressed data).
//
// Implement a rot13Reader that implements io.Reader
// and reads from an io.Reader, modifying the stream
// by applying the rot13 substitution cipher to all alphabetical characters.
//
// The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.
type rot13Reader struct {
	r io.Reader
}

func Rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'M':
		b = b + 13
	case 'M' < b && b <= 'Z':
		b = b - 13
	case 'a' <= b && b <= 'm':
		b = b + 13
	case 'm' < b && b <= 'z':
		b = b - 13
	}
	return b
}

func (rot *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot.r.Read(b)
	// fmt.Printf("n = %v, err = %v\n", n, err)
	// fmt.Printf("b=%v\n", b)
	if err == io.EOF {
		return
	}
	for i, value := range b {
		b[i] = Rot13(value)
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	n, err := io.Copy(os.Stdout, &r)
	fmt.Printf("\nresult: n=%v, err=%v", n, err)
}
