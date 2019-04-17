package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	// Calling a method on an interface value executes the method of the same name on its underlying type.
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	// Under the hood, interface values can be thought of as a tuple of a value and a concrete type
	fmt.Printf("(value: %v, type: %T)\n", i, i)
}
