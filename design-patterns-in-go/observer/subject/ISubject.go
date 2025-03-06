package subject

import "github.com/prateek69/observerPattern/subscriber"

type ISubject interface {
	AddSubscriber(subscriber subscriber.ISubscriber)
	RemoveSubscriber(subscriber subscriber.ISubscriber)
	notifySubscribers()
}
