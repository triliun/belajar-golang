package main

import "fmt"

/*
Secara default di Go-Lang semua variable di passing by value, bukan by reference
Artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lainnya,
sebenarnya yang di kirim adalah duplikasi value nya (jadi duplikat dulu baru dikiri datanya)
*/
type Address struct {
	City, Province, Country string
}

func main() {
	// Contoh tidak menggunakan pointer
	// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1

	// address2.City = "Bandung"

	// fmt.Println(address1) // datanya tidak berubah
	// fmt.Println(address2) // baru berubah menjadi bandung field City nya

	/*
	   sekarang gimana saya tidak mau pass by value saya maunya pass by reference itu?
	   Pointer adalah kemampuan membuat refenrence ke lokasi data di memori yang sama, tanpa
	   menduplikasikan data yang sudah ada

	   Sederhananya dengan kemampuan pointer kita bisa membuat pass by reference

	   Gimana cara membuatnya?
	   kita harus tau Operator &
	   Untuk membuat sebuah variable dengna nilai pointer ke variable yang lain, kita bisa menggunakan
	   operator & diikuti dengan nama variable nya
	*/

	// ini cara manual tanpai alias :=
	// var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	// var address2 *Address = &address1 // pointer ke address1

	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 // pointer

	address2.City = "Bandung"

	fmt.Println(address1) // ikut berubah karena address2 pointer ke address1
	fmt.Println(address2) // berubah menjadi bandung field City nya
}
