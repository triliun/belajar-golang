package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai64) // melebihi kapasitas maka akan balik lagi ke belakang (balik lagi ke paling bawah)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Varrel Al Jabaar K"

	var e uint8 = name[0] // jadi name[] tipedatanya uint8
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)

}
