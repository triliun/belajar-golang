package main

import (
	"belajar-golang-dasar/database"   // function init otomatis dieksekusi ketika packagenya digunakan
	_ "belajar-golang-dasar/internal" // gunalan blank (_) indentifier untuk eksekusi tanpa menggunakannya
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
