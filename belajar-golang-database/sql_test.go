package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

/*
# Eksekusi Perintah SQL

Saat membuat aplikasi menggunakan database, sudah pasti kita ingin berkomunikasi dengan database menggunakan SQL
Di Go-Lang juga menyediakan function yang bisa kita gunakan untuk mengirim perintah SQL ke databse menggunakan function
(DB)ExecContext(context, sql, params(opsional variadic argument))

Ketika megnirim perintah SQL, kita butuh mengirimkan context, dan seperti yang sudah pernah kita pelajari di course Go-Lang
Context, dengan context, kita bisa mengirim sinyal cancel jika kita ingin membatalkan pengiriman perintah SQL nya
*/
func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO costumer(id, name) VALUES ('joko','Joko')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success import new costumer")
}

/*
# Query Sql

Untuk operasi SQL yang tidak membutuhkan hasil, kita bisa menggunakan hasil, kita bisa menggunakan perintah Exec, namun
jika kita membutuhkan result, seperti SELECT SQL, kita bisa menggunakan function yang berbeda

Function untuk melakukan query ke database, bisa menggunaknak function (DB)QueryContext(context, sql, params)

*/

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id , name FROM costumer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	/*
		# Rows

		Hasil Query Function adalah sebuah data structs sql.Rows
		Rows digunakan untuk melakukan iterasi terhadap hasil dari query

		Kita bisa menggunakan function (Rows)Next(boolean) untuk melakukan iterasi terhapap data hasil query, jika return
		data false, artinya sudah tidak ada data lagi didalam result

		Untuk membaca tiap data, kita bisa menggunakan (Rows) Scan(columns...)
		Dan jangan lupa, setelah menggunakan Rows, jangan lupa untuk menutupnya menggunakan (Rows) Close()
	*/

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name) // gunakan poitner untuk menangkap datanya

		if err != nil {
			panic(err)
		}

		fmt.Println("ID : ", id)
		fmt.Println("Name : ", name)
		fmt.Println("---------") // hanya untuk pemisah

	}

}

/*
# Tipe Data Column

Sebelumnya kita hanya membuat table dengan tipe data di kolom nya berupa VARCHAR
Untuk VARCHAR di database, biasanya kita gunakan String di Go-Lang
Bagaimana dengan tipe data yang lain?
Apa representasi di Go-Lang, misal tipe data timestamp, data dan lain-lain

# Mapping Tipe Data di Database

VARCHAR, CHAR = string
INT,BIGINT = int32, int64
FLOAT,DOUBLE = float32, float64
BOOLEAN = bool
DATA, DATETIME, TIME, TIMESTAMP = time.Time


# Nullable Type

Golang Database tidak mengerti tipe data NULL di database
Oleh karena itu, khusus yang bisa NULL di database, akan jadi masalah jika kita
melakukan Scan secara bulat-bulat menggunakan tipe data representasinya di Golang

# Error Data NULL
Konversi secara otomatis NULL tidak didukung oleh Driver MySql Go-Lang
Oleh karena itu, khusus tipe kolom yang bisa NULL, kita perlu menggunakan tipe data yang ada di dalam pakcage sql

string = database/sql.NullString
bool = database/sql.NullBool
float64 = database/sql.NullFloat64
int32 = database/sql.NullInt32
int64 = database/sql.NullInt64
time.Time = database/sql.NullTime
*/

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id , name, email, balance, rating, birth_date, married, created_at FROM costumer"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birth_date sql.NullTime
		var created_at time.Time
		var married bool
		err = rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &created_at) // gunakan poitner untuk menangkap datanya

		if err != nil {
			panic(err)
		}

		fmt.Println("==========") // hanya untuk pemisah
		fmt.Println("ID : ", id)
		fmt.Println("Name : ", name)
		if email.Valid {
			fmt.Println("Email : ", email.String)
		}

		fmt.Println("Balance : ", balance)
		fmt.Println("Rating : ", rating)
		if birth_date.Valid {
			fmt.Println("Birt Date : ", birth_date.Time)
		}
		fmt.Println("Married : ", married)
		fmt.Println("Created At : ", created_at)

	}
}

/*
#Sql Dengan Parameter

Saar membuat aplikasi, kita tidak mungkin melakukan hardcode perintah SQL di kode Go-Lang kita
Biasanya kita akan menerima input data dari user, lalu membuat perintah SQL dari input user,
dan mengirimnya menggunakan perintah SQL

# SQL Injection

SQL Injection adalah sebuah teknik yang menyalahgunakan sebuah celah keamanaan yang terjadi
dalam lapisan basis data sebuah aplikasi

Biasanya, SQL Injection dilakukan dengan mengirim input data dari user dengan perintah yang salah,
sehingga menyebabkan hasil SQL yang kita buat menjadi tidak valid

SQL Injection sangat berbahaya, jika kita salah membuat SQL, bisa jadi data kita tidak aman
Solusinya?
 - Jangan membuat query SQL secara manual dengan menggabungkan String secara bulat-bulat
 - Jika kita membutuhkan parameter ketika membuat SQL, kita bisa menggunakan
 	function Execute atau Query
*/

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// anggap aja ini inputan dari user
	username := "admin'; #"
	password := "admin"

	// Simulasi Query Yang Rawan SQL INJEC
	script := "SELECT username FROM user WHERE username = '" + username + "' AND password='" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses login", username)
	} else {
		fmt.Println("Gagal login")

	}

}

/*
SQL Dengan Parameter
Sekarang kita sudah tahu bahayanya SQL Injection jika menggabungkan string ketika membuat query
Jika ada kebutuhan seperti itu, sebenarnya function Exec dan Query memiliki parameter tambahan yang bisa kita gunakan untuk
mensubtitusi parameter dari function tersebut ke SQL query yang kita buat

Untuk menandai sebuah SQL membutuhkan parameter, kita bisa gunakan karakter ? (tanda tanya)
Conoth :
SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1
INSERT INTO user(username, password) VALUES (?,?)
Dan lain...
*/

func TestSqlInjectionSave(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// anggap aja ini inputan dari user
	username := "admin'; #"
	password := "admin"

	// Simulasi Query Yang Rawan SQL INJEC
	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses login", username)
	} else {
		fmt.Println("Gagal login")

	}

}

// ini adalah penggunaan Exec dengan aman
func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "Varrel; DROP TABLE user; #"
	password := "Varrel"

	script := "INSERT INTO user(username, password) VALUES (?,?)"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success import new user")
}

/*
# Auto Increment

Kadang kita membuat sebuah table dengan id auto increment
Dan kadang pula, kita ingin mengambil data id yang sudah kita insert ke dalam MySql
Sebenarnya kita bisa melakukan query ulang ke database menggunakan SELECT LAST_INSERT_ID()

Tapi untungnya di Go-Lang ada cara yang lebih mudah
Kita bisa menggunakan function (Result)LastInsertId() untuk mendapatkan id terakhir yang dibuat secara auto increment
Result adalah object yang dikembalikan ketika kita menggunakan function Exec
*/

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "varrel@gmail.com"
	comment := "Test Comment"

	script := "INSERT INTO comments(email, comment) VALUES (?,?)"
	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	// tampilkan id auto increment
	fmt.Println("Success import new comment with id", insertId)
}

/*
# Query atau Exec dengan Parameter

Saat kita menggunakan Function Query atau Exec yang menggunakan parameter, sebenarnya implementasi dibawah nya menggunakan
Prepare Statement

Jadi tahapan pertama statement nya disiapkan terlebih dahulu, setelah itu baru di isi dengan parameter
Kadang ada kasus kita ingin melakukan beberapa hal yang sama sekaligus, hanya berbeda parameternya.
Misal insert data langsung banyak
Pembuatan Prepare Statement bisa dilakukan dengan manual, tanpa harus menggunakan Query atau Exec dengan parameter

# Prepare Statement

Saat kita membuat Prepare Statement, secara otomatis akan mengenali koneksi database yang digunakan
Sehingga ketika kita mengeksekusi Prepare Statement berkali-kali, maka akan menggunakan koneksi yang sama dan lebih
effesien karena pembuatan nya hanya sekali diawal salah

Jika menggunakan Query dan Exec dengan parameter, kita tidak bisa menjamin bahwa koneksi yang digunakan akan sama,
oleh karena itu, bisa jadi prepare statement akan selalu dibuat berkali-kali walaupun kita menggunakna SQL yang sama

Untuk membuat Prepare Statement, kita bisa menggunakan function (DB)Prepare(ctx, sql)
Prepare Statement direpresentasikan dalam struct database/sql.Stmt
Sama seperti resource sql lainnya, Stmt harus di Close() jika tidak digunakan lagi
*/
func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments(email, comment) VALUES (?,?)"

	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "varrel" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke : " + strconv.Itoa(i)

		// Tidak perlu mamsukan query SQL nya karena kita sudah membuat prepare statementnya
		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment Id : ", id)
	}
}

/*
Database Transaction

Salah satu fitur andalan di database adalah Transaction
Secara default, semua perintag SQL yang kita kirim menggunakan GO-Lang akan otomatis di commit, atau istilahnya auto commit
Namun kita bisa menggunakan fitur transaction sehingga SQL yang kita kirim tidak secara otomatis di commit ke database

Untuk memulai transaksi, kita bisa menggunakan function (DB)Begin(), dimana akan menghasilkan struct Tx
yang merupakan representasi Transaction

Struct Tx ini yang kita gunakan sebagai pengganti DB untuk melakukan transaksi, dimana hampir semua function DB ada di Tx,
seperti Exec, Query, atau Prepare

Setelah selesai proses Transaction, kita bisa gunakan function (Tx)Commit() untuk melakukan commit
atau (Tx)Rollback() untuk membatalkan commit
*/

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments(email, comment) VALUES (?,?)"

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// Contoh Menggunakna Prepare Stament Pada tx
	statement, err := tx.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	// do transaction

	for i := 0; i < 10; i++ {
		email := "varrel" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke : " + strconv.Itoa(i)

		// Tidak perlu mamsukan query SQL nya karena kita sudah membuat prepare statementnya
		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment Id : ", id)
	}

	// Commit data
	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	// Menggagalkan Commit
	// err = tx.Rollback()
	// if err != nil {
	// 	panic(err)
	// }

}
