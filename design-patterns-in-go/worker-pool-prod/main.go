package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	N := 1000
	tasks := make([]Task, N)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		start := time.Now()
		SimpleExecution(tasks)
		end := time.Since(start)
		fmt.Printf("Time taken to process tasks normally is %d ms \n", end.Milliseconds())
	}()

	go func() {
		defer wg.Done()
		start := time.Now()
		concurrency := 5

		workerPool := WorkerPool{}
		workerPool.Init(concurrency)
		workerPool.Process(tasks)

		end := time.Since(start)
		fmt.Printf("Time taken to process tasks concurrently is %d ms \n", end.Milliseconds())
	}()

	wg.Wait()
}
