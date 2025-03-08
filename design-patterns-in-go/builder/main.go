package main

import (
	"fmt"

	"github.com/prateek69/builderPattern/builder"
)

func main() {
	fmt.Println("User builder!")

	director := builder.GetNewDirector()
	user1 := director.GetUserWithNameAndCity("John", "NYU")
	fmt.Printf("User name is: %s \n", user1.Name)
	fmt.Printf("User city is: %s \n", user1.City)

	user2 := director.GetUserWithName("Sam")
	fmt.Printf("User name is: %s \n", user2.Name)
	fmt.Printf("User city is: %s \n", user2.City)
}
