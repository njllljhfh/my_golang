package main

import "fmt"

// 一个字符串是 Unicode 编码的字符（或称之为 rune）集合，因此您也可以用它迭代字符串。
// 每个 rune 字符和索引在 for-range 循环中是一一对应的。它能够自动根据 UTF-8 规则识别 Unicode 编码的字符。
// 运行代码，从输出结果中我们可以看到，常用英文字符使用 1 个字节表示，而汉字使用 3 个字符表示。
func main() {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for pos, char := range str {
		fmt.Printf("Character on position %d is: %c \n", pos, char)
	}
	fmt.Println()

	str2 := "Chinese: 中国人"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for pos, char := range str2 {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}
	fmt.Println()

	fmt.Println("index    int(rune)    rune    char bytes")
	for index, rune_ := range str2 {
		fmt.Printf("%-2d      %d      %U '%c' % X\n", index, rune_, rune_, rune_, []byte(string(rune_)))
	}
}
