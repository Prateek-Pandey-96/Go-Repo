package main

import (
	"fmt"
	"time"
)

func NormalProcessing(nums []int) {
	fmt.Println("Normal Processing")

	for i := range len(nums) {
		nums[i] *= 2
		time.Sleep(1 * time.Microsecond)
	}

	for i := range len(nums) {
		nums[i] += 10
		time.Sleep(1 * time.Microsecond)
	}

	for i := range len(nums) {
		nums[i] %= 3
		time.Sleep(1 * time.Microsecond)
	}
}
