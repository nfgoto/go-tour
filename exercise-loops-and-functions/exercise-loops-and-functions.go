package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)

	for z < 10 {
		z -= (z*z - x) / (2 * z)
		z++
	}

	return z

}

func main() {
	fmt.Println(Sqrt(2))
}
