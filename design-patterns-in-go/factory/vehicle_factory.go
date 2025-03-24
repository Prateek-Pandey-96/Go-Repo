package main

import "github.com/prateek69/factoryPattern/vehicle"

type VehicleFactory struct{}

func GetVehicleFactory() *VehicleFactory {
	return &VehicleFactory{}
}

func (vf *VehicleFactory) getVehicleInstance(vehicleString string) vehicle.IVehicle {
	var vehicleInstance vehicle.IVehicle = nil
	if vehicleString == "Bus" {
		vehicleInstance = vehicle.GetBus()
	} else if vehicleString == "Bike" {
		vehicleInstance = vehicle.GetBike()
	} else {
		vehicleInstance = vehicle.GetTruck()
	}
	return vehicleInstance
}
