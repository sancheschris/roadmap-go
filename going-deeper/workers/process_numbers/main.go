package main

import (
	"fmt"
	"sync"
)

// Process a list of numbers

// Idea: Each worker receives a number and calculates, for example, its square.

// Task: Create 3 workers that process 10 numbers and return the results in order.


type Job struct {
	ID int
	Num int
}

type Result struct {
	JobID int
	Value int
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	for job := range jobs {
		fmt.Printf("Worker %d processing number %d\n", id, job.Num)
		output := job.Num * job.Num
		results <- Result{JobID: job.ID, Value: output}
	}
	defer wg.Done()
}

func main() {
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	var wg sync.WaitGroup

	// start workers
	for w := 0; w < numWorkers ; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// send jobs
	for j := 0; j < numJobs; j++ {
		jobs <- Job{ID: j, Num: j}
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	outputs := make(map[int]int)
	for res := range results {
		outputs[res.JobID] = res.Value
	}

	for i := 0; i < numJobs; i++ {
		fmt.Printf("Number %d -> square = %d\n", i, outputs[i])
	}
}