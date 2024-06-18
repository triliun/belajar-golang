package entity

/*
# Repository Pattern

Dalam buku Domain-Driven Design, Eric Evans menjelakan bahwa "Repository is a mechanism for encapsulating storage
, retrieval, and search behavior, wich emulates a collection of objects"

Patern Repository ini biasanya digunakan sebagai jembatan antar business logic aplikasi kita dengan semua
perintah SQL ke database

Jadi semua perintah SQL akam ditulis di Repository, sedangkan business logic kode program kita
hanya cukup menggunakan Repository tersebut


# Entity / Model

Dalam pemograman berorientasi object, biasanya sebuah table di database akan selalu dibuat representasinya sebagai
class Entity atau Model, namun di Go-Lang karena tidak mengenal Class, jadi kita akan representasikan data dalam
bentuk struct

Ini bisa mempermudah kita membuat kode program
Misal ketika kita query ke Repository, dibanding mengembalikan array, alangkah baiknya Repository melakukan konversi
terlebih dahulu ke struct Entity / Model, sehingga kita tinggal menggunakan obejectnya saja.
*/


type Comment struct {
	Id      int32
	Email   string
	Comment string
}


