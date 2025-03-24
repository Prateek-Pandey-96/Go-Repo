package main

import (
	"fmt"
	"sync"

	"github.com/prateek69/factoryPattern/vehicle"
)

func main() {
	fmt.Println("Lets implement factory pattern")

	vehicleFactory := GetVehicleFactory()
	wg := &sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		var vehicleInstance vehicle.IVehicle = vehicleFactory.getVehicleInstance("Bus")
		vehicleInstance.StartVehicle()
	}()

	go func() {
		defer wg.Done()
		var vehicleInstance vehicle.IVehicle = vehicleFactory.getVehicleInstance("Bike")
		vehicleInstance.StartVehicle()
	}()

	go func() {
		defer wg.Done()
		var vehicleInstance vehicle.IVehicle = vehicleFactory.getVehicleInstance("Truck")
		vehicleInstance.StartVehicle()
	}()

	wg.Wait()
}
