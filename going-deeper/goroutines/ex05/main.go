package main

import (
	"fmt"
	"time"
)


func main() {
	go func() {
		fmt.Println("Hello from anonymous Goroutine")
	}()

	time.Sleep(500 * time.Millisecond)
}