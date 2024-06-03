package main

import "fmt"

/*
Interface Kosong

Go-Lang bukanlah bahasa pemograman yang berorientasi objek

Biasanya dalam pemograman berorientasi objek, ada satu data parent di puncak yang bisa
dianggap sebagai semua implementasi data yang ada di bahasa pemograman tersebut contohnya di java ada java.lang.Object

Untuk menangani kasus seperti ini, di Go-Lang kita bisa menggunakan interface kosong

Interface kosong adalah interface yang tidak memiliki deklarasi method apapun, hal ini membuat
secara otomatis semua tipe data akan di implementasikan

Interface kosong, juga memiliki type alias bernama any
*/

func Ups() any { // any adalah alias dari interface{}
	// return 1
	// return true
	return "Ups"
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}
