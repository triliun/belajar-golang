package main

import "fmt"

/*
Interface adalah tipe data Abstract, dia tidak memiliki implementasi langsung
Sebuah interface berisikan definisi-definisi method, dan
biasanya interface digunakan sebagai kontrak
*/

type HasName interface {
	GetName() string
}

func SayHello(hasname HasName) {
	fmt.Println("Hello", hasname.GetName())
}

/*
Implementasi Interface

Setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut

Sehingga kita tidak perlu mengimplementasikan interface secara manual

Hal ini agak berbeda dengan bahasa pemograman lain yang ketika membuat interface, kita harus
menyebutkan secara explisit akan menggunakan interface yang mana
*/

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{"Varrel"}
	SayHello(person)

	animal := Animal{"Katze"}
	SayHello(animal)

}
