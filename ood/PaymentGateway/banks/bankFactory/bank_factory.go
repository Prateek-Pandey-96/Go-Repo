package bankfactory

import (
	"fmt"

	"github.com/prateek96/paymentGateway/banks/bankStrategy"
	"github.com/prateek96/paymentGateway/entities"
	"github.com/prateek96/paymentGateway/payments/paymentStrategy"
)

type BankFactory struct{}

func GetBankFactory() *BankFactory {
	return &BankFactory{}
}

func (bf *BankFactory) GetBankInstance(bank entities.BANKNAME, paymentStrategy paymentStrategy.IPaymentStrategy) bankStrategy.IBank {
	if bank == entities.BOI {
		fmt.Println("Bank used is BOI")
		return &bankStrategy.BoiBank{Name: "Bank of India", PaymentStrategy: paymentStrategy}
	} else if bank == entities.HDFC {
		fmt.Println("Bank used is HDFC")
		return &bankStrategy.HdfcBank{Name: "HDFC bank", PaymentStrategy: paymentStrategy}
	} else if bank == entities.ICICI {
		fmt.Println("Bank used is ICICI")
		return &bankStrategy.IciciBank{Name: "ICICI bank", PaymentStrategy: paymentStrategy}
	} else {
		fmt.Println("Bank used is SBI")
		return &bankStrategy.SbiBank{Name: "State bank of India", PaymentStrategy: paymentStrategy}
	}
}
