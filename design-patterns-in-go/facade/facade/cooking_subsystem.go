package facade

import "fmt"

type CookingSubsystem struct{}

func GetCookingSubsystem() *CookingSubsystem {
	return &CookingSubsystem{}
}

func (os *CookingSubsystem) CookFood() {
	fmt.Println("Order is being prepared!")
}
