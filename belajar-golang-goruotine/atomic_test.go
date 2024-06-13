package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

/*
Atomic

Go-Lang memiliki package yang bernama sync/atomic
Atomic merupakan package yang digunakan untuk menggunakan data primitive secara aman pada proses concurrent

Contohnya sebelumnya kita telah menggunakan Mutex untuk melakukan locking ketika ingin menaikan angka di counter, Hal ini sebernarnya
bisa digunakan menggunakan Atomic package
*/

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		group.Add(1)
		go func() {
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter = ", x)
}
