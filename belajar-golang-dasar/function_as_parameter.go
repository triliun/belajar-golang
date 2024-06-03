package main

import "fmt"

/*
Function Type Declaration

kadang jika nama function terlalu pangjang, agak riber untuk menuliskannya di parameter

Type Declaration juga bisa digunakan untuk membuat alias function, sehingga akan
mempermudah kita menggunakan function sebagai parameter
*/
type Filter func(string) string

// func sayHelloWithFilter(name string, filter func(string) string)
func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

/*
	Function as Parameter

	function tidak hanya bisa kita simpang di dalam variabel sebagai value

	Namun juga bisa kita gunakan sebagai parameter untuk function lain
*/
func main() {
	sayHelloWithFilter("Varrel", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)

}
