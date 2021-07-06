package main

import "fmt"

func main() {
	p := new([]int)
	fmt.Printf("%v", *p == nil)
}
