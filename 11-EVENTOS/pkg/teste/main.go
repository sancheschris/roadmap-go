package main

import "fmt"

func main() {
	evento := []string{"test", "test2", "test3", "test4"}
	
	evento = append(evento[:0], evento[2:]...)

	fmt.Println(evento)
}