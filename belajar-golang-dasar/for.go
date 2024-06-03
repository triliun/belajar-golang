package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan Ke ", counter)
		counter++
	}
	fmt.Println("Selesei")

	/*
		Dalam For, kita dapat menambahkan statement, dimana terdapat 2 statement yang bisa ditambahkan

		init statement yaitu statement sebelum dieksekusi

		post statement yaitu statement yang akan selalu dieksekusi di akhir perulangan
	*/
	for init := 1; init <= 10; init++ {
		fmt.Println("Perulangan Ke ", init)
	}
	fmt.Println("Selesei")

	// For Range bisa digunakan untuk melakukan iterasi terhadap semua data colection
	// contoh data collection : Array, Slice dan Map

	names := []string{"Varrel", "Al", "Jabaar"}

	// Akses Slice dengan Manual
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// Akses Slice dengna For Range
	// for index, name := range names {
	// 	fmt.Println("Index", index, "Name", name)
	// }

	for _, name := range names { // _ digunakan jika kita tidak ingin menggunakan Variabel nya
		fmt.Println(name)
	}
}
