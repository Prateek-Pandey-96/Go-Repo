package type1

import (
	"fmt"
	"sync"
)

type Singelton struct{}

var singeltonInstance *Singelton = nil
var mu *sync.Mutex = &sync.Mutex{}

func GetSingelton() *Singelton {
	if singeltonInstance == nil {
		mu.Lock()
		defer mu.Unlock()
		if singeltonInstance == nil {
			fmt.Println("Instance created!")
			singeltonInstance = &Singelton{}
		} else {
			fmt.Println("Instance already exists, multiple threads entered!")
		}
	} else {
		fmt.Println("Instance already exists!")
	}
	return singeltonInstance
}
