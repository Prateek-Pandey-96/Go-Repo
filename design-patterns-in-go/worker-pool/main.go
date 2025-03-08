package main

import (
	"fmt"
)

func main() {
	fmt.Println("Worker Pool pattern")
	// tweak this N to a million to see actual difference
	// also make sure to comment out print statements in concurrent execution
	N := 1000000
	// ExecuteSequentially(N)
	ExecuteConcurrently(100, N)
}
