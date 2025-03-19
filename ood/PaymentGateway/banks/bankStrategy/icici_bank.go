package bankStrategy

import (
	"github.com/prateek96/paymentGateway/entities"
	"github.com/prateek96/paymentGateway/payments/paymentStrategy"
)

type IciciBank struct {
	Name            entities.BANKNAME
	PaymentStrategy paymentStrategy.IPaymentStrategy
}

func (ib *IciciBank) Pay(amount int) bool {
	ib.PaymentStrategy.MakePayment(amount)
	return true
}
