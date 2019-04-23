package main

import (
	"fmt"
	"math"
)

const delta = 1e-6

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	z := x

	if x > 0 {
		n := 0.0
		for math.Abs(n-z) > delta {
			n, z = z, z-(z*z-x)/(2*z)
		}
		return z, nil
	}

	return 0, ErrNegativeSqrt(z)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
