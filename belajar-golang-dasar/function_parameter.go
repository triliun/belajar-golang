package main

import "fmt"

/*
Saat membuat function kadang kita membutuhkan data dari luar atau kita sebut Parameter.
Kita bisa menambahkan parameter di function bisa lebih dari satu

Parameter tidaklah wajib, jadi kita bisa membuat function tanpa parameter seperti sebelumnya yang sudah kita buat

Namun jika kita menambahkan parameter di function, maka kita wajib memasukan data ke parameternya
*/

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
func main() {
	sayHelloTo("Varrel", "Al Jabaar") // Wajib memasukan data parameter yang sudah dibuat
	sayHelloTo("Budi", "Nugraha")
}
