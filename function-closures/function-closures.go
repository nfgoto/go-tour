/*
closure is a function value that references variables from outside its body
*/
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	// returns a closure, the func is bound to the sum variable
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// Each closure is bound to its own sum variable.
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
