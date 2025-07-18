package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from Goroutine!")
}

func main() {
	go sayHello()
	fmt.Println("Hello from main function")

	time.Sleep(1 * time.Second) // Wait so the goroutine has time to execute
}

// The go sayHello() starts the function in a new goroutine.

// without time.Sleep, the main function might finish before sayHello() runs,
// since goroutines run asynchronously

