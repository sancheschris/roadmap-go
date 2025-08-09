package main

import "fmt"

// Thread 1
func main() {
	channel := make(chan string)

	// Thread 2
	go func() {
		channel <- "Hello, World"
	}()

	// Thread 1
	msg := <-channel
	fmt.Println(msg)
}