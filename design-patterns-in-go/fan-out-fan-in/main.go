package main

import "fmt"

func main() {
	fmt.Println("Let's try to find out count of people having salary more than 100k!")
	// simpleExecution()
	// pipelinedExecution()
	fanOutFanInExecution()
}

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
