package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	strLs := strings.Fields(s)
	m := map[string]int{}
	// m := make(map[string]int)

	for _, v := range strLs {
		count, ok := m[v]
		if ok {
			m[v] += count
		} else {
			m[v] = 1
		}
	}

	fmt.Println(m)
	return m

}

func main() {
	// wc.Test(WordCount)
	a := "I am learning Go! I am learning Go!"
	WordCount(a)
}
