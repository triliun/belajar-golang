// Tipe data Map
package main

import "fmt"

/*
	Pada Array atau Slice untuk mengakses data, kita menggunakan index Number yang dimulai dari 0

	Map adalah jenis tipe data yang berisikan kumpulan data yang sama,
	namun kita bisa menentukan jenis tipe data index yang kita gunakan

	Sedehananya Map adlaah tipe data kumpulan key-value (kata kunci - nilai), dimana kata kuncinya
	bersifat unik, tidak boleh sama

	Berbeda dengan Array atau Slice, jumlah data yang kita masukan ke dalama Map boleh sebanyak-banyaknya,
	asalkan kata kuncinya berbeda, jika kita gunakan kata kunci yang sama,
	maka secara otomatis data sebelumnya akan diganti dengan data baru
*/
func main() {
	// var person map[stirng]string = map[stirng]string{}
	// person["name"] = "Varrel"
	// person["address"] = "Bandung"

	person := map[string]string{
		"name":    "Varrel",
		"address": "Bandung",
	}

	fmt.Println(person["Name"])
	fmt.Println(person["address"])
	fmt.Println(person)
	// fmt.Println(person["salah"]) Maka secara otomatis akan menggunkan default value string yaitu ""

	// Function Map

	// book := map[string]string{}
	book := make(map[string]string) // Membuat tipe data Map baru menggunakan make(map(type)type)
	book["title"] = "Buku Golang"
	book["author"] = "Varrel"
	book["ups"] = "salah"

	fmt.Println(book)
	fmt.Println(book["title"]) // Mengambil data Map dengan Key
	fmt.Println(len(book))     // Untuk mendapatkan jumlah data di Map

	delete(book, "ups") // Menghapus data di Map dengan Key

	fmt.Println(book)
	fmt.Println(len(book))
}
