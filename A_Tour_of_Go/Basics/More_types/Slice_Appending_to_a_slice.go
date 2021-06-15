package main

import "fmt"

// The first parameter `s` of `append` is a slice of type `T`,
// and the rest are `T` values to append to the slice.
//
// The resulting value of `append` is a slice
// containing all the elements of the original slice plus the provided values.
//
// If the backing array of `s` is too small to fit all the given values a bigger array will be allocated.
// The returned slice will point to the newly allocated array.
func main() {
	var s []int
	fmt.Println(s)
	printSlice2(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice2(s)
	fmt.Printf("&s[0]=%p, %d\n", &s[0], s[0])

	// The slice grows as needed.
	s = append(s, 1)
	printSlice2(s)
	fmt.Printf("&s[0]=%p, %d\n", &s[0], s[0])

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice2(s)
	fmt.Printf("&s[0]=%p, %d\n", &s[0], s[0])

	// 注意：上面3个 s[0] 的内存地址都不一样，说明这3个 slice 的底层 array 的内存地址不同。
}

func printSlice2(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
