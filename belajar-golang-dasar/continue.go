package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 { // ketika hasil i dibagi 2 sama dengan 0 atau genap maka akan lanjut ke post statement / iterasi / perulangan selanjutnya
			continue
		}
		fmt.Println("Perulangan ke", i)
	}
}
