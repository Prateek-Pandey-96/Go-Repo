package subscriber

import "fmt"

type EmailSubscriber struct {
}

func (es *EmailSubscriber) Notify(price int) {
	fmt.Printf("Email: Price has changed to %d \n", price)
}
