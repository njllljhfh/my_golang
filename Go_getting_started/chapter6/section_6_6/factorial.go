package main

import (
    "fmt"
)

// 6.6练习题2，3。
func main() {
    num := 10
    printNum(&num)
    fmt.Println("= = = = = = = = = = = = = = = = = = = = = = = = = = = = = = ")

    fmt.Println(factorial(5))
}

func printNum(n *int) {

    if *n > 0 {
        fmt.Println(*n)
        *n--
        printNum(n)
    } else {
        return
    }
}

// 递归计算阶乘
func factorial(i int) int {

    if i <= 1 {
        return 1
    }
    return i * factorial(i-1)
}
