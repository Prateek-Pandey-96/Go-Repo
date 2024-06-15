package main

type PersonBuilder struct {
	person Person
}

func (b PersonBuilder) SetAge(age int) PersonBuilder {
	b.person.Age = age
	return b
}

func (b PersonBuilder) SetName(name string) PersonBuilder {
	b.person.Name = name
	return b
}

func (b PersonBuilder) SetGender(gender string) PersonBuilder {
	b.person.Gender = gender
	return b
}

func (b PersonBuilder) Build() Person {
	return b.person
}
