package main

import (
	"fmt"

	"example.com/packages/matematica"
	"github.com/google/uuid"
)

func main() {
	s := matematica.Soma(10, 20)

	// result := matematica.dividir(10.5, 2.5) 

	fmt.Printf("Resultado: %v\n", s)
	// fmt.Println(result)
	fmt.Println(uuid.New())
}