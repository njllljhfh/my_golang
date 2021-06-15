package main

import "fmt"

func main() {
	m := make(map[string]int)

	// Insert or update an element in map `m`:
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	// Retrieve an element:
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// Delete an element:
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// Test that a `key` is present with a two-value assignment:
	// elem, ok = m[key]
	//
	// If `key` is in `m`, `ok` is `true`. If not, `ok` is `false`.
	//
	// If `key` is not in the map, then `elem` is the zero value for the map's element type.
	//
	// Note: If `elem` or `ok` have not yet been declared you could use a short declaration form:
	// elem, ok := m[key]
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
