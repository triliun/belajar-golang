package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("hello/world.go"))             // untuk mengambil Direktori nya = hello
	fmt.Println(filepath.Base("hello/world.go"))            // untuk mengambil Base nya = world.go
	fmt.Println(filepath.Ext("hello/world.go"))             // untuk mengambil Extention = .go
	fmt.Println(filepath.IsAbs("hello/world.go"))           // untuk mengecek apakah url / path di sini Absolute, ngambil paling depan nya C: atau drive D:
	fmt.Println(filepath.IsLocal("hello/world.go"))         // untuk mengecek apakah url / path nya Relativ
	fmt.Println(filepath.Join("hello", "world", "main.go")) // untuk menggabungkan

}
