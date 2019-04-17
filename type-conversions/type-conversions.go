package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4

	// T(v) converts the value v to the type T
	//  in Go assignment between items of different type requires an explicit conversion

	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	fmt.Println(x, y, z)
}
