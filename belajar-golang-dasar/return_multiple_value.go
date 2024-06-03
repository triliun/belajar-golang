package main

import "fmt"

/*
Function tidak hanya dapat mengembalikan satu value, tapi juga bisa multiple value

Untuk memberitahu jika function mengembalikan multipe value, kita harus menulis semua tipe data return value nya di function
*/

func getFullName() (string, string) { // () (string, string) adalah tipe data untuk return multiple value
	return "Varrel", "Al Jabaar"
}

func main() {
	// firstName, lastName := getFullName()
	// fmt.Println(firstName, lastName)

	firstName, _ := getFullName() // Gunakan _ untuk menghiraukan variable
	fmt.Println(firstName)
}
