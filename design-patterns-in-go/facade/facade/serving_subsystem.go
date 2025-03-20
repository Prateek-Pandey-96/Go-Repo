package facade

import "fmt"

type ServingSubsystem struct{}

func GetServingSubsystem() *ServingSubsystem {
	return &ServingSubsystem{}
}

func (os *ServingSubsystem) ServeOrder() {
	fmt.Println("Order has been served!")
}
