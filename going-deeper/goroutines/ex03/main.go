package main

import (
	"fmt"
	"time"
)

func count(label string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%s: %d\n", label, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go count("Goroutine")
	count("Main")
}