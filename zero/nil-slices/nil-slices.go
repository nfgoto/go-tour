package main

import "fmt"

func main() {
	// The zero value of a slice is nil.
	// nil slice: length = 0, capacity = 0
	// has no underlying array.

	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
