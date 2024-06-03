package main

import "fmt"

/*
Asterisk Operator atau Operator *

Saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut
Semua variable yang mengacu ke data yang sama tidak akan berubah
Jika kita ingin mengubah seluruh variable yang mengacu ke data tersebut, kita bisa menggunakan operator *
*/

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1
	address2.City = "Bandung"

	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah

	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"} address1 tidak berubah ketika address2 di ubah
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"} // address1 berubah ketika address2 di ubah
	fmt.Println(address1)
	fmt.Println(address2)

}
