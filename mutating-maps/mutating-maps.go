package main

import "fmt"

func main() {
	// initialize an empty map
	m := make(map[string]int)

	// Insert an element in map
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	// update an element in map
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// delete element from map
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// Test that a key is present in map (comma OK expression)
	// when not in map, v = zero value for map value type
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
