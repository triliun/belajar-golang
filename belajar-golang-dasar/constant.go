package main

import "fmt"

func main() {

	// tidak maslah jika tidak digunakan dan tidak bisa di ubah
	// const firstName string = "Varrel"
	// const lastName = "Al Jabaar"

	const (
		firstName = "Varrel"
		lastName  = "Al Jabaar"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	// error
	// firstName = "Tidak bisa di ubah"
	// lastName = "Tidak bisa di ubah"
}
