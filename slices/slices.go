package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// slice is dynamically-sized

	// The type []T is a slice with elements of type T
	// array[lowBound, highBound]	- highBound excluded
	var s []int = primes[1:4]
	fmt.Println(s)
}
