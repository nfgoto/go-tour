// Slices are like references to arrays

package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	// A slice does not store any data,
	// it just describes a section of an underlying array
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// Changing elements of a slice modifies the corresponding elements of its underlying array.
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
