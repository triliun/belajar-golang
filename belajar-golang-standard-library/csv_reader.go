package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "Varrel Al Jabaar\n" +
		"Budi Pratama Nugraha\n" +
		"Joko Moro Diah"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF { // EOF = End of File yang berati sudah ada di bari paling akhir
			break
		}

		fmt.Println(record)
	}
}
