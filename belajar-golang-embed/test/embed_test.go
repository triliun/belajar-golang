package test

/*
# Embed Package
Sejak Go-Lang versi 1.16, terdapat package baru dengan nama embed
Package Embed adalah fitur baru untuk mempermudah membaca isi file pada saat compile time secara otomatis
dimasukkan isi file nya dalam variable

# Cara Embed File
Untuk melakukan embed file ke variable, kita bisa mengimport package embed terlebih dahulu
Selanjutnya kita bisa tambahkan komentar //go:embed (nama filenya), diatas variable yang kita tuju

Variable yang dituju tersebut nanti secara otomatis akan berisi konten file yang kita inginkan secara
otomatis ketika kode Go-Lang di compile
Variable yang dituju tidak bisa disimpan di dalam function

# Embed File ke String
Embed File bisa kita lakukan ke variable dengan tipe data string
Secara otomatis isi file akan dibaca sebagai text dan masukkan ke variable tersebut
*/

import (
	"embed"
	_ "embed" // Import package embed
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string // Harus taruh di luar function

//go:embed version.txt
var version2 string

func TestString(t *testing.T) {
	fmt.Println(version)
	fmt.Println(version2)
}

/*
# Embed File ke []byte
Selain tipe data String, embed file juga bisa dilakukan ke variable tipe data []byte
Ini cocok sekali jika kita ingin melakukan embed file dalam bentuk binary, seperti gambar dan lain-lain
*/

//go:embed logo.png
var logo []byte

func TestByteArray(t *testing.T) {
	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

/*
# Embed Multiple Files
Kadang ada kebutuhan kita ingin melakukan embed beberapa file sekaligus
Hal ini juga bisa dilakukan menggunakan embed package
Kita bisa menambahkan komentar //go:embed lebih dari satu baris
Selain itu variable nya bisa kita gunakan tipe data embed.FS
*/

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}

/*
# Path Matcher
Selain manual satu per satu, kita bisa menggunakan path matcher untuk membaca multiple file yang kita inginkan
Ini sangat cocok ketika misal kita punya pola jenis file yang kita inginkan untuk kita baca
Caranya, kita perlu menggunakan path seperti pada package function path.Match
*/

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())

			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println("Content", string(content))
		}
	}
}
