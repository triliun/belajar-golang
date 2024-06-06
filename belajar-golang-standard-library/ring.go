package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

/*
Package Container/ring
adalah implementasi struktur data circular list (Seperti cincin jadi berputar terus)
Circular list adalah struktur data ring, dimana diakhir element akan kemnali ke elemenet awal (Head)
*/

func main() {
	var data *ring.Ring = ring.New(5) // 5 kapasitas

	// Menambahkan data dengan for loop
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	// Menambahkan data secara menual
	// data.Value = "Value 1"

	// data = data.Next()
	// data.Value = "Value 2"

	// data = data.Next()
	// data.Value = "Value 3"

	// data = data.Next()
	// data.Value = "Value 4"

	// data = data.Next()
	// data.Value = "Value 5"

	data.Do(func(value any) {
		fmt.Println(value)
	})
}
