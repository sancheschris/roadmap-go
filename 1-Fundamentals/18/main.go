package main

import (
	"fmt"

	"example.com/packages/matematica"
)

func main() {
	s := matematica.Soma(10, 20)

	fmt.Printf("Resultado: %v", s)
}