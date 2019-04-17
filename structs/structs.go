// struct is a collection of fields

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	// print the struct instance
	fmt.Println(Vertex{1, 2})
}
