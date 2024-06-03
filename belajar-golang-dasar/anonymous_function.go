package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

/*
Sebelumnya setiap membuat function, kita akan selalu memeberikan sebuah nama pada function tersebut

Namun kadang ada kalanya lebih mudah membuat function secara langsung di variabel atau
parameter tanpa harus membuat function terlebih dahulu

Hal tersebut dimanakan anonymous function atau function tanpa nama
*/

func main() {
	blacklist := func(name string) bool {
		return name == "Anjing"
	}
	registerUser("Varrel", blacklist)

	// bisa juga langsung masukan anonymous function seperti ini
	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})
}
