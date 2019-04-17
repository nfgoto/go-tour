package main

import "fmt"

func main() {
	// built-in make function
	// this is how to create dynamically-sized arrays
	//  allocates a zeroed array and returns a slice that refers to it
	a := make([]int, 5)
	printSlice("a", a)

	// specify a capacity as third argument
	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
