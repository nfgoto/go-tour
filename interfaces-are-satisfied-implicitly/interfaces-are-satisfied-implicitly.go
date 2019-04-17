package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	// A type implements an interface by implementing its methods.
	// There is no explicit declaration of intent, no "implements" keyword.
	var i I = T{"hello"}
	i.M()
}
