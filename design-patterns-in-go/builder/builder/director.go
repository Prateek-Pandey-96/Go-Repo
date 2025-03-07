package builder

type Director struct{}

func GetNewDirector() *Director {
	return &Director{}
}

func (d *Director) GetUserWithNameAndCity(name string, city string) User {
	userBuilder := GetNewBuilder()
	return userBuilder.
		SetName(name).
		SetCity(city).
		Build()
}

func (d *Director) GetUserWithName(name string) User {
	userBuilder := GetNewBuilder()
	return userBuilder.
		SetName(name).
		Build()
}
