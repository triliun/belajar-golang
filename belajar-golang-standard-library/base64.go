package main

import (
	"encoding/base64"
	"fmt"
)

/*
Package Encoding
Go-Lang menyediakan package encoding untuk melakukan encode dan decode
Go-Lang menyediakan berbagai macam algoritma untuk encoding, contoh yang populer adalah
base64. CSV, XML dan json
*/

func main() {
	value := "Varrel Al Jabaar"

	var encoded = base64.StdEncoding.EncodeToString([]byte(value)) // Encode variable value
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded) // Decode varible encode
	if err == nil {
		fmt.Println(decoded)
		fmt.Println(string(decoded))
	} else {
		fmt.Println("Error", err.Error())

	}

}
