package main

import (
	"fmt"
	"slices"
)

/*
Package Slices
di Go-Lang versi terbaru, terdapat fitur bernama Generic Programming, fitur ini akan kita bahas khusus di kelas Golang Generic

Fitur Generic ini membuat kita bisa membuat parameter dengna tipe data yang bisa berubah-ubah,
tanpa harus menggunakan interface kosong atau any

Salah satu package yang menggunakan fitu generic ini adalah package slices
package slice ini digunakan untuk memanipulasi data di slice
*/

func main() {
	names := []string{"John", "Paul", "George", "Ringo"}
	values := []int{100, 95, 80, 90}

	fmt.Println(slices.Min(names)) // Mencari nilai terendah
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names)) // Mencari nilai tertinggi
	fmt.Println(slices.Max(values))

	fmt.Println(slices.Contains(names, "Varrel")) // mencari berdasarkan kata kunci
	fmt.Println(slices.Index(names, "Varrel"))    // mencari index
	fmt.Println(slices.Index(names, "Paul"))
}
