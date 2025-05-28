package main

import (
	"fmt"
)

func main() {
	// Create a slice with make() function

	slice := make([]string, 5)

	slice[0] = "Hello"
	slice[1] = "World"
	slice[2] = "from"
	slice[3] = "Go"
	slice[4] = "slices"
	fmt.Println("Slice:", slice)

	slice = append(slice, "test")

	fmt.Println(slice)
}