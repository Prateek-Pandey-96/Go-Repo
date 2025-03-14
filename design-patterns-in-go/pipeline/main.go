package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Pipeline pattern")
	// create a slice to be processed
	N := 10000
	nums := make([]int, N)
	for i := range N {
		nums[i] = i
	}

	// Operations
	// 1> multiply each number by 2
	// 2> add 10 to it
	// 3> mod by 3
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		start := time.Now()
		Pipelining(nums)
		end := time.Since(start)
		fmt.Printf("Time taken to process numbers under pipelining is %d ms \n", end.Milliseconds())
	}()

	go func() {
		defer wg.Done()
		start := time.Now()
		NormalProcessing(nums)
		end := time.Since(start)
		fmt.Printf("Time taken to process numbers normally is %d ms \n", end.Milliseconds())
	}()

	wg.Wait()
}
