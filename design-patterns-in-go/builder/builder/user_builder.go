package builder

type UserBuilder struct {
	user User
}

func GetNewBuilder() *UserBuilder {
	return &UserBuilder{}
}

func (ub *UserBuilder) SetName(name string) *UserBuilder {
	ub.user.Name = name
	return ub
}

func (ub *UserBuilder) SetCity(city string) *UserBuilder {
	ub.user.City = city
	return ub
}

func (ub *UserBuilder) Build() User {
	return ub.user
}
