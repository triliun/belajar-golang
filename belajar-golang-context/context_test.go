package belajar_golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

/*
Membuat Context
Karena Context adalah sebuah interface, untuk membuat context kita butuh sebuah struct yang sesuai dengan kontrak interface Context
Namun kita tidak perlu membuatnya secara manual
Di Go-Lang Context terdapat function yang bisa kita gunakan untuk membuat Context

----------

Function Membuat Context
context.Background()
Untuk membuat context kosong, Tidak pernah dibatalkan, tidak pernah timeout, dan tidak memiliki value apapun.
Biasanya digunakan di main function atau dalam test, atau dalam awal proses request terjadi

context.TODO()
Membuat context kosong seperti Background(), namun biasanya menggunakan ketika belum jelas context apa yang ingin digunakan
*/

func TestContext(t *testing.T) {
	backgorund := context.Background()
	fmt.Println(backgorund)

	todo := context.TODO()
	fmt.Println(todo)
}

/*
Parent dan Child Context

Context menganut konsep parent dan child
Artinya, saat kita membuat context, kita bisa membuat child context dari context yang sudah ada / yang kita buat

Parent context bisa memiliki banyak child, namun child hanya bisa memiliki satu parent context
Konsep ini mirip dengan pewarisan di pemograman berorientasi object


----------

Hubungan Antara Parent dan Child Context
Parent dan Child context akan selalu terhubung
Saat nanti kita melakukan misal pembatalan context A, maka semua child dan sub child dari context A akan ikut dibatalkan

Namun jika misal kita membatalkan context B, hanya context B dan semua child dan sub child nya yang akan dibatalkan,
parent context B tidak akan ikut dibatalkan

Begitu juga nanti saat kita menyisipkan data ke dalam context A, semua child dan sub child nya bisa mendapatkan data tersebut

Namun jika kita menyisipkan data di context B, hanya context B dan semua child dan sub child dan sub child nya yang mendapat data,
parent context B tidak akan mendapat data.

----------

Immutable
Context merupakan object yang immutable, artinya setelah Contect dibuat, dia tidak bisa diubah lagi

Ketika kita menambahkan value ke dalam context, atau menambahkan data atau merubah pengaturan timaout dan yang lainnya,
secara otomatis akan membenuk child contect baru, bukan context tersebut
*/

/*
Context With Value
Pada saat awal membuat context, context tidak memiliki value

Kita bisa menambahkan sebuah value dengan data Pair (key - value) ke dalam context

Saat ktia menambahkan sebuah value ke context, secara otomatis akan tercipta child context baru, artinya
original context tidak akan berubah sama sekali

Untuk membuat menambahkan value ke context, kita bisa menggunakan function context.WithValue(parent, key, value)
*/

func TestContextWithValue(t *testing.T) {
	// Membuat Context kosong
	contextA := context.Background()

	// Membuat context child B dan C dari context A (Parent)
	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	// Membuat context child D dan E dari context B (Parent)
	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	// Membuat context child F dari context C (Parent)
	contextF := context.WithValue(contextC, "f", "F")

	// Membuat context child G dari context F (Parent)
	contextG := context.WithValue(contextF, "g", "G")

	// Print hasil nya
	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	// Mengambil Value context
	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c")) // dapat dari parent
	fmt.Println(contextF.Value("b")) // tidak dapat, karena beda parent
	fmt.Println(contextA.Value("b")) // tidak bisa mengambil data child nya karena ngambilnya ke atas(ngambilnya ke parent bukan ke child) bukan ke bawah
}

/*
Context With Cancel
Selain menambahkan value ke context, kita juga bisa menambahkan sinyal ke cancel ke context

Kapan sinyal cancel di perlukan dalam context?
Biasanya ketika kita butuh menjalankan proses lain, dan kita ingin bisa memberi sinyal cancel ke proses tersebut

Biasanya proses ini berupa goroutine yang berbeda, sehingga dengna mudah jika kita ingin membatalkan eksekusi goroutine,
kita bisa mengirim sinyal cancel ke context nya

Namun ingat, goroutine yang menggunakan context, tetap harus melakukan pengecekan terhadap context nya,
jika tidak, tidak ada gunanya

Untuk membuat context dengan signal, kita bisa menggunakan function context.WithCancel(parent)
*/

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			// kalau sudah selesei kita return
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // simulate slow respond
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent) // signal untuk cancel

	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}
	cancel() // Mengirim signal cancel ke context supaya goroutine nya berhenti dan tidak terjadi goroutine leak

	time.Sleep(5 * time.Second)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}

/*
Context With Timeout
Selain menambahkan value ke context, dan juga sinyal ke cancel, kita juga bisa menambahkan sinyal cancel ke context secara
otomatis dengan menggunakan pengaturan timeout

Dengan menggunakan pengaturan timeout, kita tidak perlu melakukan eksekusi cancel secara manual, cancel akan otomatis di
eksekusi jika waktu timeout sudah terlewati

Penggunaan context dengan timeout sangat cocok ketika misal melakukan query ke database atau http api, namun ingin
menentukan batas maksimal timeout nya

Untuk membuat context dengan cancel signal secara otomatis menggunakan timeout, kita bisa menggunakan
function context.WithTimeout(pareng, duration)
*/

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}

/*
Context With Deadline
Selain menggunakan timeout untuk melakukan cancel secara otomatis, kita juga bisa menggunakan deadline

Pengaturan deadline sedikti berbeda dengan timeout, jika timeout kita beri waktu dari sekarang,
kalau deadline kita bisa tentukan kapan waktu timeoutnya, misal jam 12 siang hari ini

Untuk membuat context dengan cancel signal secara otomatis menggunakan deadline, kita bisa menggunkan
function context.WithDeadline(parent, time)
*/

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}
