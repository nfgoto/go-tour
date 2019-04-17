package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// methods are just functions with a receiver
// here a function that performs the same as the method
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
