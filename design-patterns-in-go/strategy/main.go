package main

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/prateek69/strategyPattern/strategy"
)

func main() {
	fmt.Println("Using strategy pattern to calculate parking charges")

	fmt.Println("Enter hours: ")
	var input string
	fmt.Scanln(&input)

	hours, _ := strconv.Atoi(input)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		var normalStrategy strategy.IStrategy = &strategy.NormalStrategy{}
		parkingService := ParkingService{costCalculationStrategy: normalStrategy}
		fmt.Printf("Total charges using normal strategy are %d Rupees \n", parkingService.GetCost(hours))
	}()

	go func() {
		defer wg.Done()
		var peakHoursStrategy strategy.IStrategy = &strategy.PeakHoursStrategy{}
		parkingService := ParkingService{costCalculationStrategy: peakHoursStrategy}
		fmt.Printf("Total charges using peak hour strategy are %d Rupees \n", parkingService.GetCost(hours))
	}()
	wg.Wait()

}
