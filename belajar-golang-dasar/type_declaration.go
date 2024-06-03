package main

import "fmt"

func main() {
	type NoKTP string // Membuat alias type data

	var ktpVarrel NoKTP = "1111111111"

	var contoh string = "2222222222"
	var contohKtp NoKTP = NoKTP(contoh) // bisa di conversion juga seperti biasa

	fmt.Println(ktpVarrel)
	fmt.Println(contohKtp)
}
