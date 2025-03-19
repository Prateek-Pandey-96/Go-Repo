package paymentStrategy

import "fmt"

type NetBanking struct {
	UserName *string
	Password *string
}

func GetNetBankingPayment(username *string, password *string) IPaymentStrategy {
	return &NetBanking{
		UserName: username,
		Password: password,
	}
}

func (nb *NetBanking) MakePayment(amount int) (bool, error) {
	if nb.UserName == nil || nb.Password == nil {
		return false, nil
	}
	fmt.Printf("Payment made of %d through net banking.\n", amount)
	return true, nil
}
