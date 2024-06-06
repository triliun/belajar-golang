package main

/*
Package Container/List
adalah implementasi struktur data double linked list di Go-Lang
*/

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New() // Membuat Container/List

	// Memasukan data pada list
	data.PushBack("Varrel")
	data.PushBack("Al")
	data.PushBack("Jabaar")

	// Memanggil Container/List secara manual
	var head *list.Element = data.Front() // Mengambil data dari depan
	fmt.Println(head.Value)               // Varrel

	next := head.Next() // Al
	fmt.Println(next.Value)

	next = next.Next() // Jabaar
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() { // iterasi menggunakan for loop
		fmt.Println(e.Value)
	}
}
