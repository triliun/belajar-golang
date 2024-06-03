package main

import "fmt"

/*
Function as Value

Function adalah first class citizen (kalau di Golang) artinya kita bisa menyimpan function di sebuah variabel atau dijadikan sebuah value

jadi function merupakan tipe data dan bisa di simpan di variabel
*/

func getGoodBye(name string) string {
	return "Goodbay" + name
}

func main() {
	// contoh := getGoodBye() jika menggunakan kurung buka - kurung tutup artinya kita memanggil function nya
	contoh := getGoodBye // menjadikan function sebagai variabel
	misal := getGoodBye

	fmt.Println(contoh("Varrel"))
	fmt.Println(misal("Varrel Al Jabaar"))
}
