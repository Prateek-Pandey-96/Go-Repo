package main

import (
	"fmt"
	"time"
)

func ExecuteSequentially(N int) {
	fmt.Println("Sequential execution!")
	start := time.Now()

	for range N {
		time.Sleep(1 * time.Microsecond)
	}

	end := time.Since(start)
	fmt.Printf("Program took %d milliseconds \n", end.Milliseconds())

}
