package main

import (
	"fmt"
)

// the code finish before the goroutine func1 finishes

func main() {
	go func1()
	func2()
}

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println("func1:", i)
	}
}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("func2:", i)
	}
}