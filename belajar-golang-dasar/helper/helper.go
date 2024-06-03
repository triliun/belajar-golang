package helper

import "fmt"

/*
Package adalah tempat yang bisa digunakan untuk mengorganisir kode program yang kita buat di Go-Lang
Dengan menggunakan package, kita bisa merapikan kode program yang kita buat
Package sendiri sebenarnya hanya direktor folder di sistem operasi yang kita gunakan

seblumnya kita menggunakan nama package main dan function nya main juga, karena itu adalah
ketentuan dari Go-Lang untuk menjalankan kode program yang kita buat

Kalian juga bisa membuat package sendiri
nama package mengikuti folder yang kita buat
*/

/*
Bagaimana caranya menggunakan function di file lainnya?

Secara standar, file Go-Lang hanya bisa mengakses file Go-Lang lainnya yang berada dalam
package yang sama

Jika kita mengakses file Go-Lang yang berada diluar package, maka kita bisa menggunakan import
*/

/*
Access Modifier

Di bahasa pemograman lain, biasanya ada kata kunci yang bisa digunakan untuk menentukan access modifier
terhadap suatu function atau variable

Di Go-Lang, untuk menentukan access modifier, kita cukup ubah dengan nama function atau variable

Jika nama nya diawali dengan huruf besar maka artinya bisa diakses dari package lain, jika dimulai
dengan huruf kecil artinya tidak bisa diakses dari package lain tetapi bisa diaksesdi package yang sama
*/

var version = "1.0.0"
var Application = "Golang"

func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}

func Contoh() {
	sayGoodBye("Varrel")
	fmt.Println(version)
}
