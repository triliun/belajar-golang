// Tipe data Slice
package main

import "fmt"

/*
Slice adalah potongan dari data Array
Slice mirip dengan Array, bedanya ukuran Slice bisa berubah-rubah / flexibel
Slice dan Array selalu terkoneksi jadi Slice yang mengelola sebagian / seluruh Array
*/
func main() {
	/*
		tipe data Slice memiliki 3 data, yaitu Pointer, Length dan Capacity
		Pointer adalah penunjuk data pertama pada Slice
		Length adalah panjang data dari Slice
		Capacity adalah kapasitas dari Slice, dimana Length tidak boleh lebih dari Capacity

		array[low:high] Membuat Slice dari array dimulai index low sampai index sebelum high
		array[low:] Membuat Slice dari array dimulai index low sampai index akhir high
		array[:high] Membuat Slice dari array dimulai index 0 sampai index sebelum high
		array[:] Membuat Slice dari array dimulai index 0 sampai index akhir di array
	*/
	names := [...]string{"Vrl", "Al", "Jabaar", "Joko", "Budi", "Nugraha"}

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	// var slice4 []string = names[:] ini adalah contoh manual tipe data Slice []string
	slice4 := names[:]
	fmt.Println(slice4)

	// Function Slice
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:] // Sabtu, Minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	/*
		append([]string, slice, data) Menambahkan Slice baru dengan menambahkan data ke posisi terakhir slice,
		jika kapasitas / Capacity sudah penuh maka akan membuat Array baru
	*/
	daySlice2 := append(daySlice1, "Libur Baru")

	daySlice2[0] = "Sabtu Lama"
	/*
		Membaut Slice baru karena sudah melebihi Capacity
		daysBaru := days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu", "Libur Baru"}
	*/

	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Vrl"
	newSlice[1] = "Al"
	// newSlice[2] = "Jbr"	Error karena panjang / length yang kita tentukan adalah 2 harusnya menggunakan Append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice)) // Untuk mendapatkan panjang Slice
	fmt.Println(cap(newSlice)) // Untuk mendapatkan kapasita Slice

	newSlice2 := append(newSlice, "Jbr")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice) // Berubah karena masih menggunakan Array yang sama kecuali Capacity nya sudah tidak bisa menampung lebih dari 5

	// Copy Slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice)) // Untuk membuat Slice baru

	copy(toSlice, fromSlice) // Untuk Menyalin Slice dari Source ke Destination

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// Perbedaan Array dan Slice
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
