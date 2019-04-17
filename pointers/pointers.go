//  holds the memory address of a value
// The type *T is a pointer to a T value.
// Its zero value is nil.

package main

import "fmt"

func main() {
	// pointer declaration
	var p *int

	i, j := 42, 2701

	// point to i
	// & operator generates a pointer to its operand
	p = &i

	// read i through the pointer
	// * operator denotes the pointer's underlying value
	fmt.Println("*p (dereferenced pointer) =", *p)

	fmt.Println("p (pointer to i) =", p)

	// set i through the pointer
	*p = 21

	// see the new value of i
	fmt.Println("i =", i)

	// point to j
	p = &j

	// divide j through the pointer
	*p = *p / 37

	// see the new value of j
	fmt.Println("*j (dereferenced/indirected pointer) =", j)

	fmt.Println("p (pointer to j) =", p)
}
