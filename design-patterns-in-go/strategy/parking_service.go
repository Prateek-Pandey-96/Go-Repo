package main

import "github.com/prateek69/strategyPattern/strategy"

type ParkingService struct {
	costCalculationStrategy strategy.IStrategy
}

func (ps *ParkingService) GetCost(hours int) int {
	return ps.costCalculationStrategy.CalculateCharges(hours)
}
