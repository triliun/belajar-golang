package main

import (
	"fmt"
	"regexp"
)

/*
Package regecp atau regular expression

adalah utilitas di Go-Lang untuk melakukan pencarian menggunakan regular expression
Regualar Expression di Go-Lang menggunakan library C yang dibuat oleh Google bernara RE2
*/

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("varrel"))
	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edo"))

	fmt.Println(regex.FindAllString("varrel al jabaar budi eko joko edo eto ebo eyo", 10)) // max pencarian 10
}
