package main

import (
	"fmt"

	"github.com/prateek69/adapterPattern/adapter"
)

func main() {
	fmt.Println("Creating adapter for a currency converter!")

	var currency adapter.Currency = adapter.Currency{}
	var adapter adapter.IAdapter = &adapter.ConcreteAdapter{Adaptee: currency}

	fmt.Printf("The original converted value is %d\n", currency.Convert(5, 80))
	fmt.Printf("The converted value after using adapter is %d\n", adapter.ConvertCurrency(5, 80))
}
