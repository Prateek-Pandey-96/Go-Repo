package vehicle

import "fmt"

type Bus struct{}

func GetBus() IVehicle {
	return &Bus{}
}

func (b *Bus) StartVehicle() {
	fmt.Println("Bus Started!")
}
