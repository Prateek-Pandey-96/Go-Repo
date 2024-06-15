package main

import "fmt"

func main() {
	fmt.Println("Let's learn builder creational pattern!")
	var personBuilder = PersonBuilder{}
	person1 := personBuilder.SetAge(27).SetGender("Male").SetName("Prateek").Build()
	fmt.Println(person1)

	person2 := personBuilder.SetGender("Female").SetName("Rashmi").Build()
	fmt.Println(person2)
}
