package main

/*
Package os
Go-Lang telah menyediakan banyal sekali package bawaan, salah satunya adalah package os
Package os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen (bisa gunakan disemua sistem operasi)
*/

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args // fungsi untuk mendapatkan argument ketika menjalankan aplikasi contoh "go run os.go Varrel AL Jabaar"

	for _, args := range args {
		fmt.Println(args)
	}

	// fungsi untuk menagambil nama host nya di aplikasi yang kita gunakan
	name, err := os.Hostname()

	if err == nil {
		fmt.Println(name)
	} else {
		fmt.Println("Err", err.Error())

	}
}
