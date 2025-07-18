package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, i)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	go printMessage("A")
	go printMessage("B")

	printMessage("C")
}