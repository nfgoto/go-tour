package main

import (
	"fmt"
	"math"
)

// numeric type
type MyFloat float64

// method on a numeric type
// can only declare methods on types defined in the same package as method
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
