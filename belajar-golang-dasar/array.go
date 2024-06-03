package main

import "fmt"

// array / larik(indonesia)
func main() {
	var name [3]string // kapasitas 3 jika lebih maka terjadi Error

	name[0] = "Varrel"
	name[1] = "Al"
	name[2] = "Jabaar"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	// var values [3]int = [3]int{90,80,95} OPSIONAL
	var values = [...]int{ // datanya flexibel dapat menyesuaikan
		90,
		80,
		95,
		100,
		110,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values)) // untuk mendapatkan panjang / kapasitasnya
}
