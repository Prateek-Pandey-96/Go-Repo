package main

import (
	"fmt"
	"sync"

	"github.com/prateek69/singelton/type1"
	"github.com/prateek69/singelton/type2"
)

func main() {
	fmt.Println("Lets implement singelton using two methods!")
	fmt.Println("Type-1 We create a singelton and call it throughout the code, same instance is returned!")
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		_ = type1.GetSingelton()
	}()
	go func() {
		defer wg.Done()
		_ = type1.GetSingelton()
	}()
	wg.Wait()
	_ = type1.GetSingelton()

	fmt.Println("Type-2 We create a dependency and pass it throughout the code, same instance is used!")
	dependency := &type2.Dependency{
		SingeltonInstance: &type2.Singelton{},
	}

	wg.Add(2)
	go func(d *type2.Dependency) {
		defer wg.Done()
		_ = d.SingeltonInstance
	}(dependency)
	go func(d *type2.Dependency) {
		defer wg.Done()
		_ = d.SingeltonInstance
	}(dependency)
	wg.Wait()
	_ = dependency.SingeltonInstance

}
