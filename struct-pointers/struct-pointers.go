package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}

	// Struct fields can be accessed through a struct pointer
	p := &v

	// could be written (*p).X but Go allows to write p.X
	p.X = 1e9
	fmt.Println(v)
}
