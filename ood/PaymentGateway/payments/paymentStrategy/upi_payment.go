package paymentStrategy

import "fmt"

type UPI struct {
	UpiId *string
}

func GetUpiPayment(upiId *string) IPaymentStrategy {
	return &UPI{
		UpiId: upiId,
	}
}

func (u *UPI) MakePayment(amount int) (bool, error) {
	if u.UpiId == nil {
		return false, nil
	}
	fmt.Printf("Payment made of %d through UPI.\n", amount)
	return true, nil
}
