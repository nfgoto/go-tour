package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// range form of the for loop
	// ranging over a slice
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// used to range over slices and maps
}
