package bankStrategy

type IBank interface {
	Pay(amount int) bool
}
