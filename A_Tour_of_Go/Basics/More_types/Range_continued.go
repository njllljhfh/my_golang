package main

import "fmt"

func main() {
	// If you only want the index, you can omit the second variable.
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	// You can skip the index or value by assigning to `_`.
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
