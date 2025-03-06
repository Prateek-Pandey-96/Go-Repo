package main

import (
	"fmt"

	"github.com/prateek69/observerPattern/subject"
	"github.com/prateek69/observerPattern/subscriber"
)

func main() {
	fmt.Println("Stock price notifier!")

	stockSubject := &subject.StockSubject{}
	stockSubject.Init()

	emailSubscriber := &subscriber.EmailSubscriber{}
	messageSubscriber := &subscriber.MessageSubscriber{}

	stockSubject.AddSubscriber(emailSubscriber)
	stockSubject.AddSubscriber(messageSubscriber)

	stockSubject.ChangePrice(10)
	stockSubject.ChangePrice(10)

	stockSubject.RemoveSubscriber(emailSubscriber)
	stockSubject.ChangePrice(20)
}
