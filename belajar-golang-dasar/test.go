package main

import "fmt"

type HashName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func SayHello(name HashName) {
	fmt.Println("Hello", name.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	person := Person{"Varrel"}
	person1 := Person{Name: "Bima"}
	animal := Animal{"Kucing"}

	SayHello(person)
	SayHello(person1)
	SayHello(animal)
}
