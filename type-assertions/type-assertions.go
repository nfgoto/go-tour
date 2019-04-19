/*
	A type assertion provides access to an interface value's underlying concrete value.

	t := i.(T)

	This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t

		If i does not hold a T, the statement will trigger a panic.

*/
package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// this is a type assertion providing acces to the underlying concrete value
	s := i.(string)
	fmt.Println("type assertion of string =", s)

	// test whether the i interface value holds a string type
	// no panic for the comma ok notation if value not found
	s, ok := i.(string)
	fmt.Println(s, ok)

	// f will default to the wero value
	f, ok := i.(float64)
	fmt.Println(f, ok)

	// panic because no concrete / underlying value of type float64
	f = i.(float64)
	fmt.Println(f)
}
