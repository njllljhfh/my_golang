// package main
//
// import (
// 	"fmt"
// 	"time"
// )
//
// func longCalculation() {
// 	time.Sleep(500 * time.Millisecond)
// }
//
// func main() {
// 	timer(longCalculation)
// }
//
// func timer(f func()) {
// 	start := time.Now()
// 	f()
// 	end := time.Now()
// 	delta := end.Sub(start)
// 	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
// }

package main

import (
	"fmt"
	"time"
)

// golang一行代码计算函数运行时间
func main() {
	defer timeCost(time.Now())
	fmt.Println("start program")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("finish program")
}

func timeCost(start time.Time) {
	terminal := time.Since(start)
	fmt.Println(terminal)
}
