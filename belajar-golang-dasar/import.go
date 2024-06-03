package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Varrel")
	fmt.Println(result)
	fmt.Println(helper.Application)

	// Tidak bisa diakses kecuali diawali huruf besar
	// fmt.Println(helper.version)
	// fmt.Println(helper.sayGoodBye("Varrel"))
}
