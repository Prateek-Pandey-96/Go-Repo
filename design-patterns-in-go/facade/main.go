package main

import (
	"fmt"

	"github.com/prateek69/facadePattern/facade"
)

func main() {
	fmt.Println("Lets see a restaurant facade in action!")

	orderSubsystem, cookingSubsystem, servingSubsystem :=
		facade.GetOrderSubsystem(),
		facade.GetCookingSubsystem(),
		facade.GetServingSubsystem()

	restaurantFacade := facade.GetRestaurantFacade(
		orderSubsystem,
		cookingSubsystem,
		servingSubsystem,
	)

	restaurantFacade.GetFood()
}
