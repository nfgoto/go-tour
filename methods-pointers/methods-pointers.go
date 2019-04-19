package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// method with value receiver
// a COPY of the Vertex value is passed (typical behavior of func args)
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// method with pointer receiver
// Methods with pointer receivers can modify the value to which the receiver points
// most common form of method than value receivers (copy of the original)
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
