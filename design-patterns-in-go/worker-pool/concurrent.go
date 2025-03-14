package main

import (
	"fmt"
	"sync"
	"time"
)

type WorkerPool struct {
	wg     sync.WaitGroup
	jobs   chan int
	result chan int
}

func getNewWorkerPool() *WorkerPool {
	return &WorkerPool{
		wg:     sync.WaitGroup{},
		jobs:   make(chan int),
		result: make(chan int),
	}
}

func work(workerPool *WorkerPool) {
	defer workerPool.wg.Done()

	for i := range workerPool.jobs {
		// fmt.Printf("Processing job %d \n", i)
		time.Sleep(1 * time.Microsecond)
		workerPool.result <- i
	}

}

func ExecuteConcurrently(workers int, N int) {
	fmt.Println("Concurrent execution!")
	start := time.Now()
	workerPool := getNewWorkerPool()

	// send data
	go func() {
		defer close(workerPool.jobs)
		for i := range N {
			// fmt.Printf("Sending job %d \n", i)
			workerPool.jobs <- i
		}
	}()

	// output data
	go func() {
		defer close(workerPool.result)
		for range workerPool.result {
			// fmt.Printf("Consuming result for job %d \n", i)
		}
	}()

	// process data
	for range workers {
		workerPool.wg.Add(1)
		go work(workerPool)
	}
	workerPool.wg.Wait()
	end := time.Since(start)
	fmt.Printf("Program took %d milliseconds \n", end.Milliseconds())
}
