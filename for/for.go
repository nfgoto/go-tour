package main

import "fmt"

func main() {
	sum := 0

	// the For loop is the only looping construct in Go (1.12)

	/*
		- init statement; executed before the first iteration
		- condition expression: evaluated before every iteration
		- post statement: executed at the end of every iteration
	*/
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
