package belajar_golang_database

import (
	"database/sql"
	"time"
)

/*
# Database Pooling

sql.DB di Go-Lang sebenarnya bukanlah sebuah koneksi ke database
Melainkan sebuah pool ke database, atau dikenal dengan konsep Database Pooling *isi nya kumpulan dari koneksi

Di dalam sql.DB, Go-Lang melakukan management koneksi ke database secara otomatis. Hal ini
menjadikan kita tidak perlu melakukan management koneksi ke database secara manua

Dengan kemampuan database pooling ini, kita bisa menentukan jumlah minimal dan maksimal koneksi yang di buat oleh Go-Lang
sehingga tidak membanjiri koneksi ke database, karena biasanya ada batas maksimal koneksi yang bisa ditangani
oleh database yang kita gunakan

# Pengaturan Database Pooling
(DB) SetMaxIdleConns(number) Pengaturan berapa jumlah koneksi minimal yang dibuat
(DB) SetMaxOpenConns(number) Pengaturan berapa jumlah koneksi maximal yang dibuat
(DB) SetConnMaxIdleTime(duration) Pengaturan berapa lama koneksi yang sudah tidak digunakan akan dihapus
(DB) SetConnMaxLifetime(duration) Pengaturan berapa lama koneksi yang boleh digunakan
*/

/*
# Error Tipe Data Date
Secara default, Driver MySql untuk Go-Lang akan melakukan query tipe data DATE, DATETIME, TIMESTAMP menjadi
[]byte / []uint8. Dimana ini bisa dikonversi menjadi String, lalu di parsing menjadi time.Time

Namun hal ini merepotkan jika dilakukan secara manual, kita bisa meminta Driver MySql untuk Go-Lang secara otomatis
melakukan parsing dengan menambahkan parameter parseTime=true
*/

func GetConnection() *sql.DB {
	// Membuat database
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar-golang-database?parseTime=true")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
