package main

var a3 string

func main() {
	a3 = "G"
	print(a3)
	f1()
}

func f1() {
	a3 := "O"
	print(a3)
	f2()
}

// 注意：f2() 中的 a3 指向的是全局变量 a3
func f2() {
	print(a3)
}
