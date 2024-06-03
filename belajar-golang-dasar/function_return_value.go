package main

import "fmt"

/*
Function bisa mengembalikan data menggunkan return value
Untuk memberitahu bahwa function mengembalikan data, kita harus menuliskan tipe data pengembalian dari function tersebut

Jika function tersebut kita deklarasikan dengan tipe data pengembalian, maka wajib di dalam
function nya kita harus mengembalikan data

Untuk mengembalikan data dari function, kita bisa menggunakan kata kunci return diikuti dengan datanya
*/

func getHello(name string) string { // () string adalah tipe data untuk return
	hello := "Hello" + name
	return hello
}

func main() {
	result := getHello("Varrel")
	fmt.Println(result)

	fmt.Println(getHello("Budi"))
	fmt.Println(getHello("Bima"))
}
