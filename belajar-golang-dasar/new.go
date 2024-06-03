package main

import "fmt"

/*
Operator New

Sebelumnya kita sudah membuat pointer dengan menggunakan operator &
Go-Lang juga memiliki function new yang bisa digunakan untuk membaut pointer
Namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal
*/

type Address struct {
	City, Province, Country string
}

func main() {
	// var alamat1 *Address = &Address{}
	// var alamat1 *Address = new(Address)
	// var alamat2 *Address = alamat1

	alamat1 := new(Address)
	alamat2 := alamat1

	alamat2.Country = "Indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
