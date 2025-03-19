package paymentStrategy

import "fmt"

type CardPayment struct {
	CardNum *string
	Cvv     *string
}

func GetCardPayment(cardNum *string, cvv *string) IPaymentStrategy {
	return &CardPayment{
		CardNum: cardNum,
		Cvv:     cvv,
	}
}

func (cp *CardPayment) MakePayment(amount int) (bool, error) {
	if cp.CardNum == nil || cp.Cvv == nil {
		return false, nil
	}
	fmt.Printf("Payment made of %d through card.\n", amount)
	return true, nil
}
