package main

import (
	"fmt"
	"time"
)

// Go programs express error state with `error` values.

// The `error` type is a built-in interface similar to `fmt.Stringer`:
// type error interface {
//    Error() string
// }

// 注意：(As with `fmt.Stringer`, the `fmt` package looks for the `error` interface when printing values.)
//
// Functions often return an `error` value,
// and calling code should handle errors
// by testing whether the error equals `nil`.
//
// i, err := strconv.Atoi("42")
// if err != nil {
//    fmt.Printf("couldn't convert number: %v\n", err)
//    return
// }
//
// fmt.Println("Converted integer:", i)
// A `nil` error denotes success; a `non-nil` error denotes failure.

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	// 在Go语言中，格式化时间 "yyyy-MM-dd HH:mm:ss" 为特定的数字 "2006-01-02 15:04:05" 该时间是Go语言的创建时间.
	return fmt.Sprintf("at %v, %s",
		e.When.Format("2006-01-02 15:04:05"), e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
