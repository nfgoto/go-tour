package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for index, value := range pow {
		fmt.Printf("1 << uint(%d) = %d\n", index, value)
	}
}
