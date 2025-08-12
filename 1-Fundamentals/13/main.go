package main

import "fmt"

func soma(a, b *int) int {
	*a = 30
	return *a + *b
}

func main() {
	a := 10
	b := 15

	fmt.Println(soma(&a, &b))
}