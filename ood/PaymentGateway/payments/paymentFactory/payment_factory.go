package paymentFactory

import (
	"github.com/prateek96/paymentGateway/entities"
	"github.com/prateek96/paymentGateway/payments/paymentStrategy"
)

type PaymentFactory struct {
}

func GetPaymentFactory() *PaymentFactory {
	return &PaymentFactory{}
}

func (pf *PaymentFactory) GetPaymentModeInstance(mode entities.Mode) paymentStrategy.IPaymentStrategy {
	var strategy paymentStrategy.IPaymentStrategy
	if mode == entities.CARD {
		cardNum := "card_num"
		cvv := "cvv"
		strategy = paymentStrategy.GetCardPayment(&cardNum, &cvv)
	} else if mode == entities.NetBanking {
		username := "user_name"
		password := "password"
		strategy = paymentStrategy.GetNetBankingPayment(&username, &password)
	} else {
		upiId := "upi_id"
		strategy = paymentStrategy.GetUpiPayment(&upiId)
	}
	return strategy
}
