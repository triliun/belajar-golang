package main

import "fmt"

func main() {
	name := "Joko"

	switch name {
	case "Varrel":
		fmt.Println("Hallo Varrel")
	case "Budi":
		fmt.Println("Hallo Budi")
	default:
		fmt.Println("Hai, Boleh Kenalan?")
	}

	// Short Statement switch
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// Switch tanpa kondisi
	name = "Al Jabaar"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
