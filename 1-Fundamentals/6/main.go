package main

import "fmt"

func main() {
	s := []int{2, 4, 6, 8, 10}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// when I do an append on a slice, it double the capacity. It creates a new array under the hoods
	s = append(s, 22)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)	

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])	

	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	fmt.Printf("len=%d cap=%d %v\n", len(s[1:]), cap(s[1:]), s[1:])
}