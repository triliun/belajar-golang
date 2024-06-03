package main

import "fmt"

func random() any {
	// return "OK"
	return 1000
}

func main() {
	var result any = random()
	// var resultString string = result.(string) // konfersi interface{} ke string
	// fmt.Println(resultString)aman

	/*
		Type Assertions Menggunakan Switch

			Saat salah menggunakan type assertions, maka bisa berakibat terjadi panic di aplikasi kita
			Jika panic dan tidak ter-recover, maka otomatis program kita akan mati

			Agar lebih aman, sebaiknya kita menggunakan switch expression untuk melakukan type assertions
	*/

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unkown", value)

	}

}
