package main

import (
	"fmt"

	bankfactory "github.com/prateek96/paymentGateway/banks/bankFactory"
	"github.com/prateek96/paymentGateway/entities"
	"github.com/prateek96/paymentGateway/payments/paymentFactory"
	"github.com/prateek96/paymentGateway/storage"
)

func main() {
	fmt.Println("Payment Gateway using Strategy and factory method pattern!")

	pg := PaymentGateway{
		storage:        storage.GetInMemStorage(),
		router:         getRouter(),
		bankFactory:    *bankfactory.GetBankFactory(),
		paymentFactory: *paymentFactory.GetPaymentFactory(),
	}

	client := entities.Client{Id: "one", Name: "Flipkart"}
	// Requirement - 1 Client flow
	pg.AddClient(&client)
	// fmt.Println(pg.HasClient(&client))
	// pg.RemoveClient(&client)
	// fmt.Println(pg.HasClient(&client))

	// Requirement - 2 Paymode flow
	pg.AddPaymode("UPI")
	pg.AddPaymode("CARD")
	pg.AddPaymode("NETBANKING")
	// fmt.Println(pg.ListPaymodes())
	// pg.RemovePaymode("CARD")
	// fmt.Println(pg.ListPaymodes())

	// Requirement - 3 Make payment flow
	pg.ClientPay(&client, "CARD", 10)
	pg.ClientPay(&client, "UPI", 10)
	pg.ClientPay(&client, "NETBANKING", 10)

	// Can be extended for :-
	// 1> Additional api's in pg_service that can change router
	// 2> A counter can be implemented in case of multiple banks for a paymode to distribute requests
	// 3> A blocking call can be introduced to input values required for specific payment strategies (hardcoded currently)
}
