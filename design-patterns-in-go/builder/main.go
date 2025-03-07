package main

import (
	"fmt"

	"github.com/prateek69/builderPattern/builder"
)

func main() {
	fmt.Println("User builder!")

	user1 := builder.GetNewDirector().GetUserWithNameAndCity("Prateek", "Jaipur")
	fmt.Printf("User name is: %s \n", user1.Name)
	fmt.Printf("User city is: %s \n", user1.City)

	user2 := builder.GetNewDirector().GetUserWithName("Prateek")
	fmt.Printf("User name is: %s \n", user2.Name)
	fmt.Printf("User city is: %s \n", user2.City)
}
