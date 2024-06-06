package main

import (
	"fmt"
	"path"
)

/*
Package Path
package path digunakan untuk memanipulasi data path seperti path di URL atau path di File System

Secara default Package path menggunakan slash sebagai karakter pathnya, oleh karena itu cocok untuk data URL

Namun jika ingin menggunkan untuk memanipulasi path di File System, kare Windows mengunakan backslash,
maka khusus untuk File System, perlu menggunakan package path/filepath
*/

func main() {
	fmt.Println(path.Dir("hello/world.go"))             // untuk mengambil Direktori nya = hello
	fmt.Println(path.Base("hello/world.go"))            // untuk mengambil Base nya = world.go
	fmt.Println(path.Ext("hello/world.go"))             // untuk mengambil Extention = .go
	fmt.Println(path.Join("hello", "world", "main.go")) // untuk menggabungkan

}
