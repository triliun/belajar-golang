package main

import "fmt"

/*
Recover adalah function yang bisa kita gunakan untuk menangkap data panic

Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan
*/

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println(message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
	// jangan menaruh recover function disini, karena jika terjai panic tidak akan di eksekusi
}

func main() {
	runApp(true) // jika terjadi panic maka kode dibawah tidak akan di eksekusi
	fmt.Println("Varrel Al Jabaar K")
}
