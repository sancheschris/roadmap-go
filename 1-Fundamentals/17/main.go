package main

import "fmt"

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m1 := map[string]int{"Chris": 1000, "Gabi": 2000}
	m2 := map[string]float64{"Edna": 2500.40, "Rafa": 300.50}

	fmt.Println(m1)
	fmt.Println(m2)
}