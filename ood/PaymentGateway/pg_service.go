package main

import (
	"math/rand"

	bankfactory "github.com/prateek96/paymentGateway/banks/bankFactory"
	"github.com/prateek96/paymentGateway/entities"
	"github.com/prateek96/paymentGateway/payments/paymentFactory"
	"github.com/prateek96/paymentGateway/storage"
)

type PaymentGateway struct {
	storage        storage.IStorage
	bankFactory    bankfactory.BankFactory
	paymentFactory paymentFactory.PaymentFactory
	router         map[entities.Mode][]entities.BANKNAME
}

// client related start
func (pg *PaymentGateway) AddClient(client *entities.Client) {
	_ = pg.storage.AddClient(client)
}
func (pg *PaymentGateway) RemoveClient(client *entities.Client) {
	_ = pg.storage.RemoveClient(client)
}
func (pg *PaymentGateway) HasClient(client *entities.Client) bool {
	hasClient, _ := pg.storage.HasClient(client)
	return hasClient
}

// client related end

// paymode related start
func (pg *PaymentGateway) AddPaymode(paymode string) {
	_ = pg.storage.AddPaymode(entities.Mode(paymode))
}
func (pg *PaymentGateway) RemovePaymode(paymode string) {
	_ = pg.storage.RemovePaymode(entities.Mode(paymode))
}
func (pg *PaymentGateway) ListPaymodes() []entities.Mode {
	paymodes, _ := pg.storage.ListPaymodes()
	return paymodes
}

// paymode related end
func (pg *PaymentGateway) ClientPay(client *entities.Client,
	paymentMode string, amount int) bool {

	randomIndex := rand.Intn(len(pg.router[entities.Mode(paymentMode)]))
	bankName := pg.router[entities.Mode(paymentMode)][randomIndex]

	paymentModeInstance := pg.paymentFactory.GetPaymentModeInstance(entities.Mode(paymentMode))

	bank := pg.bankFactory.GetBankInstance(bankName, paymentModeInstance)
	bank.Pay(amount)
	return true
}
