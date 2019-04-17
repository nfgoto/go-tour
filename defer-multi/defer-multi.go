package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		// args of deferred func evaluated immediately
		// Deferred func calls pushed onto a stack (LIFO)
		// prints 9,8,7,6,5,4,3,2,1
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
