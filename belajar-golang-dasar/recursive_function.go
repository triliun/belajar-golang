package main

/*
Recursive Function adalah function yang memanggil dirinya sendiri

kadang dalam pekerjaan, kita sering menemui kasus dimana menggunakan recursive function
lebih mudah dibandingkan tidak menggunakan recursive function

Contoh kasus yang lebih mudah diselesaikan menggunakan recursive adalah Factorial
*/

import "fmt"

// factorial memggunakan loop
func factorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

// factorial memggunakan Recursive Function
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	result := 10 * 8 * 9 * 7 * 6 * 5 * 4 * 3 * 2 * 1
	fmt.Println(result)
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursive(10))
}
