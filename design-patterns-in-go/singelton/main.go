package main

import (
	"fmt"
	"sync"
)

type config struct{}

var mutex *sync.Mutex = &sync.Mutex{}
var configInstance *config = nil
var counter int = 1

func GetConfig() {
	if configInstance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if configInstance == nil {
			fmt.Println("Config instance created - ", counter)
			configInstance = &config{}
			counter = counter + 1
		} else {
			fmt.Println("Config instance already created Condition-1! Returning that one")
		}
	} else {
		fmt.Println("Config instance already created Condition-2! Returning that one")
	}
}

func main() {
	fmt.Println("Let's learn singelton creational pattern!")
	for i := 0; i < 100; i++ {
		go GetConfig()
	}
	fmt.Scanln()
}
