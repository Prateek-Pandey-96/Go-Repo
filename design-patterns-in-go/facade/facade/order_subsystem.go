package facade

import "fmt"

type OrderSubsystem struct{}

func GetOrderSubsystem() *OrderSubsystem {
	return &OrderSubsystem{}
}

func (os *OrderSubsystem) TakeOrder() {
	fmt.Println("Order has been taken!")
}
