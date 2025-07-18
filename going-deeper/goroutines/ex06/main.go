package main

import (
	"fmt"
	"sync"
)

// Synchronizing Goroutines with WaitGroup

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // signal that this goroutine is done
	fmt.Printf("Worker %d starting\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait() // Wait until all goroutines have called Done()
	fmt.Println("All workers done")
}