package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	// default low bound = 0
	s = s[:2]
	fmt.Println(s)

	// default high bound = length of slice
	s = s[1:]
	fmt.Println(s)

	// [:] => [0:length]
}
