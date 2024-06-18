package belajar_golang_database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

/*
# Membuat Koneksi Ke Database

Hal yang pertama akan kita lakukan ketika membuat aplikasi yang akan menggunakan database adalah
melakukan koneksi ke database

Untuk melakukan koneksi ke database di Go-Lang, kita bisa membuat object dari struct namanya sql.DB menggunakan
function sql.Open(driver, dataSourceName)

Untuk menggunakan database MySql, kita bisa menggunakan driver "mysql"

Sedangkan untuk dataSourceName, tiap database biasanya punya cara penulisan masing-masing misal di MySql,
kita bisa menggunakan dataSourceName seperti ini:
- username:password@tcp(host:port)/database_name

Jika object sql.DB sudah tidak di gunakan lagi, disarankan untuk menutupnya mengguankan function Close()
*/

func TestOpenConnection(t *testing.T) {
	// Membuat database
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar-golang-database")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Gunakan DB
}
