package main

/*
# Hasil Embed di Compile
Perlu diketahui, bahwa hasil embed yang dilakukan oleh package embed adalah permanent dan data file yang dibaca
disimpan dalam binary file Golang nya

Artinya bukan dilakukan secara realtime membaca file yang ada diluar
Hal ini menjadikan jika binary file golang sudah di compile, kita tidak butuh lagi file external nya,
dan bahkan jika diubah file external nya, isi variable nya tidak akan berubah lagi
*/

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed logo.png
var logo []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	// version
	fmt.Println(version)

	//logo
	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	// path
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())

			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println("Content", string(content))
		}
	}
}
