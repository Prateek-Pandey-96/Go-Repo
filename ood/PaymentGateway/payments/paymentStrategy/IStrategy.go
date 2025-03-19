package paymentStrategy

type IPaymentStrategy interface {
	MakePayment(amount int) (bool, error)
}
