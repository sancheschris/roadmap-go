package main

import "fmt"

func main() {

	// different loops

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	arr := []string{"Hello", "World", "Golang"}

	// for i, v := range arr {
	// 	fmt.Printf("Index: %d - Value: %s\n", i, v)
	// }

	for _, v := range arr {
		fmt.Println(v)
	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for {
		fmt.Println("Endless message")
	}

}