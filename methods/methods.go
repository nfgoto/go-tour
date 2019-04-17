package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// A method is a function with a special receiver argument
// The receiver appears in its own argument list
// between the func keyword and the method name (Abs).
func (thisV Vertex) Abs() float64 {
	return math.Sqrt(thisV.X*thisV.X + thisV.Y*thisV.Y)
}

func main() {
	aVertex := Vertex{3, 4}
	fmt.Println(aVertex.Abs())
}
