package main

import (
	"flag"
	"fmt"
)

/*
Package flag berisikan fungsionalitas untuk memparsing command line argument
*/

func main() {
	var host *string = flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")
	port := flag.Int("port", 0, "Put your database port")

	flag.Parse() // untuk parse flag yang kita buat

	fmt.Println("Host", *host)
	fmt.Println("Username", *username)
	fmt.Println("Password", *password)
	fmt.Println("Port", *port)

	// go run flag.go -host="12.12.12.12" -username=varrel -password="rahasia banget" -port=8080
}
