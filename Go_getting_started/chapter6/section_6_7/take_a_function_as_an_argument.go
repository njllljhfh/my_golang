package main

import (
	"fmt"
	"strings"
)

// 将函数作为参数
func main() {
	callback(1, Add)
	fmt.Println("= = = = = = = = = = = = = = = = = = = = = = = = = = = = = = ")

	fmt.Println(strings.IndexFunc("中国nb", IsAscii))
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(1, 2)
}

func IsAscii(c rune) bool {
	if c > 255 {
		return false
	}
	return true
}
