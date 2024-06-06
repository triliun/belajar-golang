package main

/*
Package Strings

adalah package yang berisikan function-funtion untuk memanipulasi tipe data String
*/
import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Varrel Al Jabaar", "Varrel"))                            // mencari kata kunci varrel balikannya bool
	fmt.Println(strings.Split("Varrel Al Jabaar", " "))                                    // memotong string berdasarkan separator
	fmt.Println(strings.ToLower("Varrel Al Jabaar"))                                       // Membuat semua karakter string menjadi lower case
	fmt.Println(strings.ToUpper("Varrel Al Jabaar"))                                       // Membuat semua karakter string menjadi uppper case
	fmt.Println(strings.Trim("              Varrel Al Jabaar           ", " "))            // Memotong cutset di awal dan di akhir string
	fmt.Println(strings.ReplaceAll("Varrel Al Jabaar Varrel Al Jabaar", "Varrel", "Budi")) // Mengubah semua string dari Varrel ke Budi
}
