package main

/*
	Duration
	Saat menggunakan tipe data waktu, kadang kita butuh data berupa durasi
	Package tipe memiliki type Duration, yang sebenarnnya adalah alias untuk int64
	Namun terdapat banyak method yang bisa kita gunakan untuk memanipulasi Duration
*/

import (
	"fmt"
	"time"
)

func main() {
	var duration1 time.Duration = time.Second * 100     // 100 detik
	var duration2 time.Duration = time.Millisecond * 10 // 10 Milidetik
	var duration3 time.Duration = duration1 - duration2

	fmt.Println(duration3)
	fmt.Printf("Duration : %d\n", duration3) // formatnya nano second
}
