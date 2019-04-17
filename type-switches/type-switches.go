/*
	A type switch is a construct that permits several type assertions in series.


*/
package main

import "fmt"

func do(i interface{}) {
	// this is a type switch
	switch v := i.(type) {
	// the cases in a type switch specify types (not values)
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
