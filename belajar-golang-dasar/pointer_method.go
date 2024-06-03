package main

/*
Walaupun method akan menempel di struct, tapi sebenarnya data di struct yang diakses di dalam
method adalah pass by value

Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena
harus selalu diduplikas ketika menggunakan method
*/

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr." + man.Name
}

func main() {
	vrl := Man{"Varrel"}
	vrl.Married()

	fmt.Println(vrl.Name)
}
