package subject

import "github.com/prateek69/observerPattern/subscriber"

type StockSubject struct {
	stockPrice  int
	subscribers map[subscriber.ISubscriber]struct{}
}

func GetNewStockSubject() *StockSubject {
	return &StockSubject{
		stockPrice:  0,
		subscribers: make(map[subscriber.ISubscriber]struct{}),
	}
}

func (ss *StockSubject) ChangePrice(price int) {
	if ss.stockPrice != price {
		ss.stockPrice = price
		ss.notifySubscribers()
	}
}

func (ss *StockSubject) AddSubscriber(subscriber subscriber.ISubscriber) {
	ss.subscribers[subscriber] = struct{}{}
}

func (ss *StockSubject) RemoveSubscriber(subscriber subscriber.ISubscriber) {
	delete(ss.subscribers, subscriber)
}

func (ss *StockSubject) notifySubscribers() {
	for subscriber := range ss.subscribers {
		subscriber.Notify(ss.stockPrice)
	}
}
