package main

import "fmt"

/*
Break & Continue adalah kata kunci yang bisa digunakan dalam perulangan
Break digunakan untuk menghentikan seluruh perulangan
Continue digunakan untuk menghentikan perulangan yang sedang berjalan dan langsung melanjutkan ke perulangan selanjutnya
*/
func main() {
	for i := 0; i < 10; i++ {
		if i == 5 { // setelah mencapai 5 maka code di bawahnya akan di hentikan
			break
		}
		fmt.Println("Perulangan ke", i)
	}
}
