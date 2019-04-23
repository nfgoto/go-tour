package main

import (
	"fmt"
	"math"
)

const delta = 1e-6

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// need to convert ErrNegativeSqrt type to the underlying float64 type to avoid infinite loop
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	z := x

	if x < 0 {
		return 0, ErrNegativeSqrt(z)
	}

	n := 0.0
	for math.Abs(n-z) > delta {
		n, z = z, z-(z*z-x)/(2*z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
