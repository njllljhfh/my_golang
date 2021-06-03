package main

import "fmt"

// A slice has both a `length` and a `capacity`.
//
// The length of a slice is the number of elements it contains.
//
// 注意: The capacity of a slice is the number of elements in the underlying array,
// counting from the first element in the slice.
//
// The length and capacity of a slice `s` can be obtained using the expressions `len(s)` and `cap(s)`.
//
// You can extend a slice's length by re-slicing it, provided it has sufficient capacity.
// Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.

// 注意：对slice，进行re-slicing时，是以底层的array为基础进行re-slicing，而不是以上一次的切片结果为基础。
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:] // 仔细观察这个切片的容量
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
