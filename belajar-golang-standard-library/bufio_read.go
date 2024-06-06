package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
Package IO
IO atau Input Output merupakan fitur di Go-Lang yang digunakan sebagai standard untuk proses Input Output

Di Go-Lang semua mekanisme intput output pasti mengukuti standard package io
*/

/*
Reader
untuk membaca input, Go-Lang menggunakan kontrak interface bernama Reader yang terdapat di package IO

Writer
untuk menulis ke output, Go-Lang menggunakan kontrak interface bernama Writer yang terdapat di package IO
*/

/*
Bagaimana Implementasi IO nya?
Pengguna dari IO sendiir di Go-Lang terdapat banyak package, sebelumnya contohnya kita menggunakan CSV Reader dan CSV Writer

Karena package IO sebenarnya hanya kontrak / interface untuk IO, untuk mengimplementasi nya kita harus lakukan sendiri

Tapi untungnya Go-Lang juga menyediakan package untuk mengimplementasi IO secara mudah, yaitu menggunakan package bufio atau buffered io
package ini digunakan untuk membuat data seperti Reader dan Writer
*/

func main() {
	input := strings.NewReader("this is long string\nwith new line\n")

	reader := bufio.NewReader(input)

	for {
		// line, prefix, err := reader.ReadLine() // Membaca per baris sampai \n
		line, _, err := reader.ReadLine() // Membaca 1 baris sampai \n
		if err == io.EOF {
			break
		}
		fmt.Println(string(line)) // variable line bertipe []byte jadi kita perlu konversi ke string

	}
}
