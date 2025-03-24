package vehicle

import "fmt"

type Truck struct{}

func GetTruck() IVehicle {
	return &Truck{}
}

func (t *Truck) StartVehicle() {
	fmt.Println("Truck Started!")
}
