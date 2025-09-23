package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for file := range jobs {
		duration := time.Duration(rand.Intn(3)+1) * time.Second
		fmt.Printf("Worker %d downloading %s...\n", id, file)
		time.Sleep(duration)
		fmt.Printf("Worker %d finished %s in %v\n", id, file, duration)
		results <- file
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	files := []string{"file1.txt", "file2.jpg", "file3.pdf", "file4.mp3", "file5.zip"}
	numWorkers := 3

	jobs := make(chan string, len(files))
	results := make(chan string, len(files))
	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for _, file := range files {
		jobs <- file
	}
	close(jobs)

	wg.Wait()
	close(results)

	downloaded := []string{}
	for file := range results {
		downloaded = append(downloaded, file)
	}

	fmt.Println("\nDownloaded files:")
	for _, file := range downloaded {
		fmt.Println(file)
	}
}