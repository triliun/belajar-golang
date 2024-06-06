package main

/*
Standard Library

Go-Lang, selain bahasa pemograman, Go-Lang juga menyediakan Standard Library (package bawaan) tanpa
harus menggunakan package dari luar buatan orang lain

Contohnya package yang sering digunakan yaitu package fmt dan errors
*/

import "fmt"

/*
Sebelumnya kita sudah sering menggunakan package fmt dengan menggunakan function Println
Selain Println, masih banyak function yang terdapat di package fmt, contohnya yang banyak digunakan untuk melakukan format
*/

func main() {
	fristName := "Varrel"
	lastName := "Al Jabaar"

	fmt.Println("Hello '", fristName, lastName, "'")

	fmt.Printf("Hello '%s %s'\n", fristName, lastName)

}
