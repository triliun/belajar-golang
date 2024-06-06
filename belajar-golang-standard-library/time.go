package main

import (
	"fmt"
	"time"
)

/*
Package Time
adalah package yang berisikan fungsionalitas untuk management waktu di Go-Lang
*/

func main() {
	now := time.Now()
	// fmt.Println(now.Local()) // untuk mendapatkan Waktu di laptop kita
	fmt.Println(now) // untuk mendapatkan Waktu di laptop kita

	var utc time.Time = time.Date(2009, time.August, 17, 0, 0, 0, 0, time.UTC) // Membuat Tanggal baru
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"

	value := "2020-10-10 10:10:10"
	valueTime, err := time.Parse(formatter, value) // untuk memparshing waktu ke string
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(valueTime)
	}

	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Hour())
}
