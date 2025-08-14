package main

import "fmt"

func main() {
	
	a := 10
	var pointer *int = &a

	fmt.Println(pointer)

	*pointer = 20
	b := &a
	fmt.Println(*b)
}