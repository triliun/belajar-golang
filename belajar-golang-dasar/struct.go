package main

import "fmt"

/*
Struct adalah sebuah template data yang di gunakan untuk menggabungkan nol atau lebih tipe data
lainnya dalam satu kesatuan

Struct biasanya representasi data dalam program aplikasi yang kita buat

Data di struct disimpan dalam field

Sederhananya struct adalah kumpulan dari field
*/

type Costumer struct { // Kebanykan orang membuat nama struct menggunakan huruh Kapital / Pascal case
	Name, Address string
	Age           int
}

/*
Struct Method

Struct  adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function

Namun jika kita ingin menambahkan Method ke dalam struct, sehingga seakan-akan sebuah struct memiliki function

Method adalah function yang menempel pada struct
*/
func (costumer Costumer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", costumer.Name)
}

func main() {
	var vrl Costumer
	fmt.Println(vrl)

	vrl.Name = "Varrel Al Jabaar"
	vrl.Address = "Bandung, Indonesia"
	vrl.Age = 18

	fmt.Println(vrl)
	fmt.Println(vrl.Name)
	fmt.Println(vrl.Address)
	fmt.Println(vrl.Age)

	// Sebelumnya kita telah membuat data dari struct, namun sebenarnya ada banyak cara yang bisa kita gunakan untuk membuat struct

	// Membuat data menggunalan nama field nya
	joko := Costumer{
		Name:    "Joko",
		Address: "Semarang, Indonesia",
		Age:     22,
	}
	fmt.Println(joko)

	// Membuat data tanpa nama field, tepai harus sesuai urutannya
	budi := Costumer{"Budi", "Jakarta, Indonesia", 16}
	fmt.Println(budi)

	// Ini adalah cara penggunaan method sayHello
	vrl.sayHello("Joko")
	joko.sayHello("Varrel")
	budi.sayHello("Joko")
}
