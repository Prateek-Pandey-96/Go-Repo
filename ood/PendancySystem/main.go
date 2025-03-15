package main

import (
	"fmt"

	"github.com/prateek96/pendancySystem/services"
	"github.com/prateek96/pendancySystem/storage"
)

func main() {
	fmt.Println("Welcome to the pendancy system")

	storage := storage.GetInMemStorage()
	pendancyService := services.GetPendancyService(storage)
	pendancySystem := PendancySystem{pendancyService: pendancyService}

	pendancySystem.StartTracking(1112, []string{"UPI", "Karnataka", "Bangalore"})
	pendancySystem.StartTracking(2451, []string{"UPI", "Karnataka", "Mysore"})
	pendancySystem.StartTracking(3421, []string{"UPI", "Rajasthan", "Jaipur"})
	pendancySystem.StartTracking(1221, []string{"Wallet", "Karnataka", "Bangalore"})

	fmt.Printf("count for UPI tag is %d \n", pendancySystem.GetCounts([]string{"UPI"}))
	fmt.Printf("count for UPI, Karnataka tags is %d \n", pendancySystem.GetCounts([]string{"UPI", "Karnataka"}))
	fmt.Printf("count for UPI, Karnataka, Bangalore tags is %d \n", pendancySystem.GetCounts([]string{"UPI", "Karnataka", "Bangalore"}))
	fmt.Printf("count for Bangalore tag is %d \n", pendancySystem.GetCounts([]string{"Bangalore"}))

	pendancySystem.StartTracking(4221, []string{"Wallet", "Karnataka", "Bangalore"})
	pendancySystem.StopTracking(1112)
	pendancySystem.StopTracking(2451)

	fmt.Printf("count for UPI tag is %d \n", pendancySystem.GetCounts([]string{"UPI"}))
	fmt.Printf("count for Wallet tag is %d \n", pendancySystem.GetCounts([]string{"Wallet"}))
	fmt.Printf("count for UPI, Karnataka, Bangalore tags is %d \n", pendancySystem.GetCounts([]string{"UPI", "Karnataka", "Bangalore"}))
}
