package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")

	// runs first case which value is equal to condition expression

	// notice short statement (os := runtime.GOOS;)
	// only runs the selected case (no cascade)
	switch os := runtime.GOOS; os {
	// cases don't need to be constants
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
