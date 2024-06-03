package main

import "fmt"

/*
Closure adalah kemampuan sebuah function berinteraksi dengan data-data disekitarnya dalam scope yang sama

Harap gunakan closure ini dengan bijak saat membuat Aplikasi
*/

func main() {
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		// kemampuan sebuah function berinteraksi dengan data data di sekitarnya, mengubah nilai variabel counter
		counter++
	}

	increment() // counter = 1
	increment() // counte = 2
	increment() // counter = 3
	fmt.Println(counter)
}
