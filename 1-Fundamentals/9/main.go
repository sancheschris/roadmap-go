package main

import "fmt"

func main() {
	fmt.Println(sum(1, 3, 45, 6, 19, 20, 35))
}

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
