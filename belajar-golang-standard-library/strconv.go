package main

/*
Package Strconv

Sebelumnya kita sudah belajar cara konversi tipe data, misa lint32 ke int64
Bagaimana jika kita butuh melakukan konversi yang tipe data nya berbeda? misal int ke string atau sebaliknya
Hal tersebut bisa kita lakukan dengan bantuan package strconv (string conversion)
*/

import (
	"fmt"
	"strconv"
)

func main() {
	resultBool, err := strconv.ParseBool("true")

	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultBool)

	}

	resultInt, err := strconv.Atoi("1000") // menggunakan PharseInt atau Atoi

	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultInt)
	}

	binary := strconv.FormatInt(999, 2) // format int ke binary
	fmt.Println(binary)

	stringInt := strconv.Itoa(999) // format int ke string
	fmt.Println(stringInt)

}
