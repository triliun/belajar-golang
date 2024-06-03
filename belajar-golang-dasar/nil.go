package main

import "fmt"

/*
NIL

Biasanya di dalam bahasa pemograman lain, objek yang belim diinisialisi maka secara otomatis
nilainya adalah null atau nil

Berbeda dengan Go-Lang, di Go-Lang saat kita buat variable dengan tipe ada tertentu, maka
secara otomatis akan dibuatkan default value nya

Namun di Go-Lang ada data nil, yaitu data kosong

Nil sendiri hanya bisa digunakan di beberapa tipde data, seperti interface, function, map, slice, pointer, dan channel
*/

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"Name": name,
		}
	}
}

func main() {
	data := NewMap("Varrel")

	if data == nil {
		fmt.Println("Data map masih kosoong")

	} else {
		fmt.Println(data["Name"])

	}
}
