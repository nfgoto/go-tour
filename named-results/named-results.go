package main

import "fmt"

func split(sum int) ( /* named returns */ x, y int) {
	x = sum * 4 / 9
	y = sum - x

	// naked return - will return all named returns
	return
}

func main() {
	fmt.Println(split(17))
}
