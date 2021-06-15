package main

import (
	"fmt"
)

// Channels
// Channels are a typed conduit through which you can send and receive values
// with the channel operator, <-.
// - - -
// ch <- v    // Send v to channel ch.
// v := <-ch  // Receive from ch, and
//           // assign value to v.
// (The data flows in the direction of the arrow.)
// - - -
//
// Like maps and slices, channels must be created before use:
// - - -
// ch := make(chan int)
// - - -
// 注意：By default, sends and receives block until the other side is ready.
// 		This allows goroutines to synchronize
// 		without explicit locks or condition variables.
//
// The example code sums the numbers in a slice,
// distributing the work between two goroutines.
// Once both goroutines have completed their computation,
// it calculates the final result.
func sum(s []int, c chan int) {
	sum1 := 0
	for _, v := range s {
		sum1 += v
	}
	// fmt.Printf("执行了么1\n")
	c <- sum1 // send sum to c
	// fmt.Printf("执行了么2\n")
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	// for i := 0; i < 100; i++ {
	// 	time.Sleep(1 * time.Second)
	// }
}
