package subscriber

import "fmt"

type MessageSubscriber struct {
}

func (es *MessageSubscriber) Notify(price int) {
	fmt.Printf("Message: Price has changed to %d \n", price)
}
