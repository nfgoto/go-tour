package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// A map maps keys to values.
// nil map
// declare a map with string keys with Vertex struct values
var m map[string]Vertex

func main() {
	// need to use make to initialize and use a nil map
	m = make(map[string]Vertex)

	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])
}
