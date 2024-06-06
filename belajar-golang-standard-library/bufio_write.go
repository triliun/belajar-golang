package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout) // output ke terminal

	_, _ = writer.WriteString("hello world\n")
	_, _ = writer.WriteString("Selamat Belajar\n")

	writer.Flush()
}
