package main

import (
	"fmt"
)

// Select
// The `select` statement lets a goroutine wait on multiple communication operations.
//
// A `select` blocks until one of its cases can run, then it executes that case.
// 注意：It chooses one at random if multiple are ready.

func fibonacci2(c, quit chan int) {
	i := 1
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
			i += 1
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	// 如果注释掉从管道中取数据的代码，则函数 fibonacci2() 中的代码 `c <- x` 会阻塞。(原因见 Channels.go 中的`注意`)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)

	// go fibonacci2(c, quit)
	// for i := 0; i < 2; i++ {
	// 	time.Sleep(1 * time.Second)
	// }
}
