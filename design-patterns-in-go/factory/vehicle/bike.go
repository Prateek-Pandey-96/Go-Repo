package vehicle

import "fmt"

type Bike struct{}

func GetBike() IVehicle {
	return &Bike{}
}

func (b *Bike) StartVehicle() {
	fmt.Println("Bike Started!")
}
