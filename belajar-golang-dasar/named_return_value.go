package main

import "fmt"

/*
Biasanya saat kita memberitahu bahwa sebuah function mengembalikan value,
maka kita hanya mendeklarasi tipe data return value fucntion

Namun kita juga bisa membuat variable secara langsung di tipe data return function nya
*/

// func getCompleteName() (string, string, string) {
// 	firstName := "Varrel"
// 	middleName := "Al"
// 	lastName := "Jabaar"
//
// return firstName, middleName, lastName
// }

// func getCompleteName() (firstName string, middleName string, lastName string) {
// 	firstName = "Varrel"
// 	middleName = "Al"
// 	lastName = "Jabaar"
//
// return firstName, middleName, lastName
// }

func getCompleteName() (firstName, middleName, lastName string) { // Variabel di return value boleh kosong
	firstName = "Varrel"
	middleName = "Al"
	lastName = "Jabaar"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
