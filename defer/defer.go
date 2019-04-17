package main

import "fmt"

func main() {
	// execute function after surrounding function returns
	// args of deferred func immediately evaluated
	// func itself executed after surrounding block
	defer fmt.Println("world")

	fmt.Println("hello")
}
