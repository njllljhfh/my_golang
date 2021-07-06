package main

import "fmt"

// 练习 6.9 不使用递归但使用闭包改写第 6.6 节中的斐波那契数列程序。
func main() {
	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
