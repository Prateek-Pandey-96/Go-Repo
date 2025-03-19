package main

import "github.com/prateek96/paymentGateway/entities"

func getRouter() map[entities.Mode][]entities.BANKNAME {
	router := make(map[entities.Mode][]entities.BANKNAME)
	router[entities.Mode("UPI")] = []entities.BANKNAME{entities.BANKNAME("SBI"), entities.BANKNAME("ICICI")}
	router[entities.Mode("CARD")] = []entities.BANKNAME{entities.BANKNAME("BOI")}
	router[entities.Mode("NETBANKING")] = []entities.BANKNAME{entities.BANKNAME("HDFC")}
	return router
}
