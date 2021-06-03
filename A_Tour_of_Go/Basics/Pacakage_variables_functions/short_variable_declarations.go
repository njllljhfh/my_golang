package main

import "fmt"

// Outside a function, every statement begins with a keyword (var, func, and so on)
// and so the `:=` construct is not available.
func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
