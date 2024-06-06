package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout) // os.Stdout membaca lewat terminal

	_ = writer.Write([]string{"varrel", "al", "jabaar"})
	_ = writer.Write([]string{"budi", "pratama", "nugraha"})
	_ = writer.Write([]string{"joko", "morro", "diah"})

	writer.Flush() // Flush() untuk mengirim semua perubahaannya
}
