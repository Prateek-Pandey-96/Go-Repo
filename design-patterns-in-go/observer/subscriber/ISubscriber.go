package subscriber

type ISubscriber interface {
	Notify(price int)
}
