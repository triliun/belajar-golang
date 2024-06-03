package main

import "fmt"

/*
Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi

Defer function akan selalu dieksekusi walaupun terjadi error di function yang di eksekusi
*/

func logging() {
	fmt.Println("Selesai memanggil function")
}

// func runAplication()  {
// 	fmt.Println("Run application")
// jikda terjadi error maka yang dibawah tidak di eksekusi
// 	logging()
// }

func runAplication() {
	defer logging() // fitur memanggil function ketika eksekusi selesei
	fmt.Println("Run application")
}

func main() {
	runAplication()
}
