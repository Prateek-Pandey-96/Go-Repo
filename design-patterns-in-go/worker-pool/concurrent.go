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
		for i := range N {
			// fmt.Printf("Sending job %d \n", i)
			workerPool.jobs <- i
		}
		close(workerPool.jobs)
	}()

	// process data
	for range workers {
		workerPool.wg.Add(1)
		go work(workerPool)
	}

	// output data
	go func() {
		for range workerPool.result {
			// fmt.Printf("Consuming result for job %d \n", i)
		}
		close(workerPool.result)
	}()

	workerPool.wg.Wait()
	end := time.Since(start)
	fmt.Printf("Program took %d milliseconds \n", end.Milliseconds())
}
