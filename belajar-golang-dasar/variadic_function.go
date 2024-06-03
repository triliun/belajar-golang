package main

import (
	"fmt"
)

/*
Parameter yang berada di posisi paling terakhir, memiliki kemampuan dijadikan sebuah varargs

Varargs artinya datanya bisa menerima lebih dari 1 input, atau anggap saja semacam Array.

Apa bedanya dengan parameter biasa dengan tipe data Array?
 - Jika parameter tipe data array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
 - Jika parameter menggunakan varargs, kita bisa langsung mengirimkan data nya, jika lebih dari 1
   cukup gunakan tanda koma
*/

func sumAll(numbers ...int) int { // value ...tipeData adalah cara membuat Variabel Argument / Varadic
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

/*
Kadang ada kasus dimana kita menggunakan Variadic Function, namun memiliki variabel berupa slice
Kita bisa menjadikan slice sebagai variadic parameter
*/

func main() {
	fmt.Println(sumAll(10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10, 10, 10, 10, 10))

	numbers := []int{10, 10, 10, 10, 10, 10}
	// fmt.Println(sumAll(numbers...)) tidak bis karena variabel numbers tipe datanya adalah slice
	fmt.Println(sumAll(numbers...)) // gunakan ... untuk conversi slice menjadi variadic function

}
