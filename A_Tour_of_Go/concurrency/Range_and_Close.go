package main

import "fmt"

// Range and Close
// A sender can `close` a channel to indicate that no more values will be sent.
// Receivers can test whether a channel has been closed
// by assigning a second parameter to the receive expression: after
// - - -
// v, ok := <-ch
// - - -
// `ok` is `false` if there are no more values to receive and the channel is closed.
//
// The loop `for i := range c` receives values from the channel repeatedly until it is closed.
//
// Note: Only the sender should close a channel, never the receiver.
// 		 Sending on a closed channel will cause a panic.
//
// Another note: Channels aren't like files; you don't usually need to close them.
// 				 Closing is only necessary when the receiver must be told
// 				 there are no more values coming, such as to terminate a `range` loop.

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}

	// dragon：测试 ok 值
	// for i := 0; i < 10*cap(c); i++ {
	// 	r, ok := <-c
	// 	if !ok {
	// 		// 管道 c 被关闭了。
	// 		fmt.Println(r, ok)
	// 		break
	// 	} else {
	// 		fmt.Println(r, ok)
	// 	}
	// }
}
