package main

/*
Pointer di Function

Saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan di
copy lalu di kirim ke function tersebut

Oleh karena itu, jika kita mengubah data di dalam function, data yang aslinya tidak akan pernah berubah.
Hal ini membuat variable menjadi aman, karena tidak akan bisa di ubah

Namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut
Untuk melakukan ini, kita juga bisa menggunakan pointer di function

Untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator asterisk / operator * di parameternya
*/

import "fmt"

/**/
type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) { // pointer menuju Address
	address.Country = "Indonesia"
}
func main() {
	var address *Address = &Address{} // gunakan tipe data pointer karena function bertipe pointer
	ChangeCountryToIndonesia(address)

	fmt.Println(address)
}
