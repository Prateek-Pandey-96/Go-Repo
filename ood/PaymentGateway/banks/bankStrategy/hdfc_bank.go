package bankStrategy

import (
	"github.com/prateek96/paymentGateway/entities"
	"github.com/prateek96/paymentGateway/payments/paymentStrategy"
)

type HdfcBank struct {
	Name            entities.BANKNAME
	PaymentStrategy paymentStrategy.IPaymentStrategy
}

func (hb *HdfcBank) Pay(amount int) bool {
	hb.PaymentStrategy.MakePayment(amount)
	return true
}
