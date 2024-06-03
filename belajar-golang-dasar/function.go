package main

import "fmt"

/*
Function adalah sebuah kode blok kode yang sengaja dibuat dalam program agar bisa digunakan berulang-rulang

Cara membuat function sangat sederhana, hanya dengan menggunakan kata kunci func lalu diikuti
dengan nama function nya dan blok kode isi function

Setelah membuat function, kita bisa mengeksekusi function tersebut dengan memanggilnya
menggunakan kata kunci nama function nya diikuti tanda kurung buka ( dan kurung tutup )
*/

func sayHello() {
	fmt.Println("Hello")
}

func main() {

	// Bebas memanggil function sesuai kebutuhan kita
	sayHello()
	sayHello()
	sayHello()
}
