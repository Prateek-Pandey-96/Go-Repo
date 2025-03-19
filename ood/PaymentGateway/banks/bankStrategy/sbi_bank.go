package bankStrategy

import (
	"github.com/prateek96/paymentGateway/entities"
	"github.com/prateek96/paymentGateway/payments/paymentStrategy"
)

type SbiBank struct {
	Name            entities.BANKNAME
	PaymentStrategy paymentStrategy.IPaymentStrategy
}

func (sb *SbiBank) Pay(amount int) bool {
	sb.PaymentStrategy.MakePayment(amount)
	return true
}
