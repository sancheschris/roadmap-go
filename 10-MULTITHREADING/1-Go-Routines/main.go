package main

import (
	"fmt"
	"time"
)

func task(task string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, task)
		time.Sleep(1 * time.Second)
	}
}

// Thread 1
func main() {
	// Thread 2
	go task("A")
	// Thread 3
	go task("B")
	go func() {
		for i := 0; i < 5; i++ {
		fmt.Printf("%d: Task %s is running\n", i, "anonymus")
		time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(15* time.Second)
}


