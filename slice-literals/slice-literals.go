package main

import "fmt"

func main() {
	// slice literal
	// creates an underlying array: [6]int{2, 3, 5, 7, 11, 13}
	// and the slice variable references it
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	// struct slice
	// creates underlying array
	/*	[6]struct {i int, b bool}{
			{2, true},
			{3, false},
			{5, true},
			{7, true},
			{11, false},
			{13, true},
		}
	*/
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
