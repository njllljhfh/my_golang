package main

import (
	"golang.org/x/tour/pic"
)

// 将函数体复制到官网教程中运行
func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)
	for i := range res {
		res[i] = make([]uint8, dx)
	}
	return res
}

func main() {
	pic.Show(Pic)
}
