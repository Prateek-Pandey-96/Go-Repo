package main

import (
	"fmt"
	"time"
)

func transform1(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range nums {
			time.Sleep(1 * time.Microsecond)
			out <- i * 2
		}
	}()
	return out
}

func transform2(nums <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range nums {
			time.Sleep(1 * time.Microsecond)
			out <- i + 10
		}
	}()
	return out
}

func transform3(nums <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range nums {
			time.Sleep(1 * time.Microsecond)
			out <- i % 3
		}
	}()
	return out
}

func Pipelining(nums []int) {
	fmt.Println("Pipelining")
	idx := 0
	for num := range transform3(transform2(transform1(nums))) {
		nums[idx] = num
		idx += 1
	}
}
