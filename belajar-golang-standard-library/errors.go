package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "Varrel" {
		return NotFoundError
	}

	// sukses
	return nil
}

func main() {
	err := GetById("Varrel")
	/*
		Mengecek Jenis Error
		Misal kita membuat jenis error sendiri, lalu kita ingin mengecek jenis errornya
		Kita bisa menggunakan errors.is() untuk mengecek jenis type error nya
	*/

	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("not found error")
		} else {
			fmt.Println("unkown error")
		}
	}
}
