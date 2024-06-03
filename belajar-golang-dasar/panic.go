package main

import "fmt"

/*
Panic function adalah function yang bisa kita gunakan untuk menghentikan program

Panic function biasanya dipanggil ketika terjadi panic pada saat program kita berjalan

Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi
*/

func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Ups Error")

	}
}

func main() {
	runApp(true)
}
