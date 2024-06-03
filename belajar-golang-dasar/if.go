package main

import "fmt"

func main() {
	name := "Varrel"

	if name == "Varrel" {
		fmt.Println("Hallo Varrel")
	} else if name == "Joko" {
		fmt.Println("Hallo Joko")
	} else if name == "Eko" {
		fmt.Println("Hallo Eko")
	} else if name == "Budi" {
		fmt.Println("Hallo Budi")
	} else {
		fmt.Println("Hai, Boleh Kenalan?")
	}

	if length := len(name); length > 5 { // Short Statement pada If
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
