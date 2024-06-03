package main

/*
Error Inteface

Go-Lang memiliki interface yang digunakan sebagai kontrak untuk membuat error, nama interface nya adalah error

Membuat Error
Untuk membuat error, kita tidak perlu melakukannya secara manual
Go-Lang sudah menyediakan library untuk membuat helper secara mudah, yang terdapat di package errors
*/
import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pemabagi int) (int, error) { // gunakan interface error
	if pemabagi == 0 {
		return 0, errors.New("Pemabagian Dengan NOL") // membuat error dengan function New()
	} else {
		return nilai / pemabagi, nil
	}
}

func main() {
	hasil, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Hasil", err.Error())

	}
}
