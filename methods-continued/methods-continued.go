package main

import (
	"fmt"
	"math"
)

// You can declare a method on NON-STRUCT types
// can only declare methods on types defined in the SAME PACKAGE as method
// to add methods on types not defined in same package, create a new type pointing to the other

// numeric type
type MyFloat float64

// method on a numeric type
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
