package main

import (
	"fmt"
	"strings"
)

// 学习并理解以下程序的工作原理：
// 一个返回值为另一个函数的函数可以被称之为工厂函数，
// 这在您需要创建一系列相似的函数的时候非常有用：
// 书写一个工厂函数而不是针对每种情况都书写一个函数。
// 下面的函数演示了如何动态返回追加后缀的函数：
func main() {
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")

	fmt.Println(addBmp("file"))  // returns: file.bmp
	fmt.Println(addJpeg("file")) // returns: file.jpeg
}

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
