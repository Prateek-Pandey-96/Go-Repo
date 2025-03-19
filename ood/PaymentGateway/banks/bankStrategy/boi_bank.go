package bankStrategy

import (
	"github.com/prateek96/paymentGateway/entities"
	"github.com/prateek96/paymentGateway/payments/paymentStrategy"
)

type BoiBank struct {
	Name            entities.BANKNAME
	PaymentStrategy paymentStrategy.IPaymentStrategy
}

func (bb *BoiBank) Pay(amount int) bool {
	bb.PaymentStrategy.MakePayment(amount)
	return true
}
